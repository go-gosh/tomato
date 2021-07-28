package interactivecli

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/go-gosh/tomato/app/ent/migrate"
	"github.com/go-gosh/tomato/app/ent/usertomato"
	"github.com/go-gosh/tomato/app/service"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"
)

const rootContainerResourceId = "root_container"

const (
	redDuration   = 25 * 60
	greenDuration = 5 * 60
)

func NewClockView(svc *service.Service, c *container.Container) (*ClockView, error) {
	clockSD, err := segmentdisplay.New()
	if err != nil {
		return nil, err
	}
	f, err := newForm()
	if err != nil {
		return nil, err
	}
	return &ClockView{
		clockSD:   clockSD,
		container: c,
		tomato:    nil,
		alarmCh:   make(chan struct{}),
		svc:       svc,
		userId:    0,
		formSD:    f,
	}, nil
}

type form struct {
	Mode   int
	Answer chan bool
	Text   *text.Text
	yesBtn *button.Button
	noBtn  *button.Button
}

func newForm() (*form, error) {
	t, err := text.New()
	if err != nil {
		return nil, err
	}
	f := &form{
		Answer: make(chan bool),
		Text:   t,
	}
	yesBtn, err := button.New("(Y)es", func() error {
		f.Answer <- true
		return nil
	}, button.GlobalKey('y'))
	if err != nil {
		return nil, err
	}
	noBtn, err := button.New("(N)o", func() error {
		f.Answer <- false
		return nil
	}, button.GlobalKey('n'))
	if err != nil {
		return nil, err
	}
	f.yesBtn = yesBtn
	f.noBtn = noBtn

	return f, nil
}

func (f *form) Layout(ctx context.Context, c *container.Container) error {
	// ask for next action
	return c.Update(rootContainerResourceId,
		container.SplitHorizontal(
			container.Top(container.PlaceWidget(f.Text)),
			container.Bottom(
				container.SplitVertical(
					container.Left(
						container.KeyFocusGroups(1, 2),
						container.PlaceWidget(f.yesBtn),
						container.AlignHorizontal(align.HorizontalRight),
						container.PaddingRight(5),
					),
					container.Right(
						container.KeyFocusGroups(1, 2),
						container.PlaceWidget(f.noBtn),
						container.AlignHorizontal(align.HorizontalLeft),
						container.PaddingLeft(5),
					),
				),
			),
		),
	)
}

type ClockView struct {
	clockSD   *segmentdisplay.SegmentDisplay
	container *container.Container
	tomato    *ent.UserTomato
	alarmCh   chan struct{}
	svc       *service.Service
	userId    int
	formSD    *form
}

func (v *ClockView) Tomato() *ent.UserTomato {
	return v.tomato
}

func (v *ClockView) SetTomato(tomato *ent.UserTomato) {
	v.tomato = tomato
}

func (v *ClockView) OnInit(ctx context.Context) error {
	tomato, err := v.svc.GetWorkingTomatoByUserId(ctx, v.userId)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	v.SetTomato(tomato)

	return v.Layout(ctx, v.container)
}

func (v *ClockView) OnSubmit(ctx context.Context) error {
	err := v.svc.CloseTomatoByUserId(ctx, 0)
	if err != nil {
		return err
	}

	if v.tomato == nil || v.tomato.Color == usertomato.ColorGreen {
		return v.OnStart(ctx)
	}

	param := service.TomatoCreate{
		Duration:  greenDuration,
		Color:     usertomato.ColorGreen,
		UserId:    0,
		StartTime: time.Now(),
	}
	t, err := v.svc.CreateTomato(ctx, param)
	if err != nil {
		return err
	}

	v.SetTomato(t)

	return nil
}

func (v *ClockView) OnStart(ctx context.Context) error {
	param := service.TomatoCreate{
		Duration:  redDuration,
		Color:     usertomato.ColorRed,
		UserId:    0,
		StartTime: time.Now(),
	}
	t, err := v.svc.CreateTomato(ctx, param)
	if err != nil {
		return err
	}

	v.SetTomato(t)

	return nil
}

func (v *ClockView) OnGiveUp(ctx context.Context) error {
	err := v.svc.GiveUpTomatoByUserId(ctx, 0)
	if err != nil {
		return err
	}

	v.SetTomato(nil)

	return nil
}

func (v *ClockView) Run(ctx context.Context) {
	err := v.OnInit(ctx)
	if err != nil {
		log.Fatalf("[ClockView.OnInit] %+v", err)
	}
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("exit")
				return
			case <-ticker.C:
				err := v.OnUpdate(ctx)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()
}

func (v *ClockView) OnUpdate(ctx context.Context) error {
	if v.tomato == nil {
		err := v.clockSD.Write([]*segmentdisplay.TextChunk{
			segmentdisplay.NewChunk(time.Now().Format("04:05")),
		})
		if err != nil {
			return err
		}

		return nil
	}

	duration := v.tomato.RemainTime.Sub(time.Now()) / time.Second
	if duration > 0 {
		spacer := " "
		if duration%2 == 0 {
			spacer = ":"
		}
		return v.clockSD.Write([]*segmentdisplay.TextChunk{
			segmentdisplay.NewChunk(
				fmt.Sprintf("%02d%s%02d", duration/60, spacer, duration%60),
			),
		})
	}
	// alarm user
	err := v.OnAlarm(ctx)
	if err != nil {
		return err
	}
	// show ask dialog
	v.formSD.Text.Reset()
	if v.tomato.Color == usertomato.ColorRed {
		err := v.formSD.Text.Write("已完成一个番茄，休息一会儿吧。")
		if err != nil {
			return err
		}
	} else {
		err := v.formSD.Text.Write("休息时间结束，开始一个新的番茄吧。")
		if err != nil {
			return err
		}
	}
	err = v.formSD.Layout(ctx, v.container)
	if err != nil {
		return err
	}
	// get user select option
	result := <-v.formSD.Answer
	if !result {
		// just close working tomato
		err := v.svc.CloseTomatoByUserId(ctx, 0)
		if err != nil {
			return err
		}
		v.SetTomato(nil)
		// re-layout this view
		err = v.Layout(ctx, v.container)
		if err != nil {
			return err
		}
		return nil
	}
	// submit tomato
	err = v.OnSubmit(ctx)
	if err != nil {
		return err
	}
	// re-layout this view
	err = v.Layout(ctx, v.container)
	if err != nil {
		return err
	}

	return nil
}

func (v *ClockView) OnAlarm(ctx context.Context) error {
	for i := 0; i < 3; i++ {
		_ = exec.Command("afplay", "/System/Library/Sounds/Tink.aiff", "-v", "15", "-q", "1").Run()
	}

	return nil
}

func (v *ClockView) Layout(ctx context.Context, c *container.Container) (err error) {
	if v.Tomato() == nil {
		btn, err := button.New("(s)tart", func() error {
			err := v.OnStart(ctx)
			if err != nil {
				return err
			}
			return v.Layout(ctx, v.container)
		}, button.GlobalKey('s'))
		if err != nil {
			return err
		}
		return c.Update(rootContainerResourceId,
			container.Border(linestyle.Light),
			container.BorderTitle("PRESS Q TO QUIT"),
			container.SplitHorizontal(
				container.Top(container.PlaceWidget(v.clockSD)),
				container.Bottom(container.PlaceWidget(btn)),
			),
		)
	}
	giveUpBtn, err := button.New("(p)ause", func() error {
		// show ask dialog
		v.formSD.Text.Reset()
		err := v.formSD.Text.Write("确定放弃当前番茄记录吗？")
		if err != nil {
			return err
		}
		err = v.formSD.Layout(ctx, v.container)
		if err != nil {
			return err
		}
		// get user select option
		result := <-v.formSD.Answer
		if result {
			// confirm give up tomato clock
			err = v.OnGiveUp(ctx)
			if err != nil {
				return err
			}
		}
		// re-layout view always
		return v.Layout(ctx, v.container)
	}, button.GlobalKey('p'))
	if err != nil {
		return err
	}
	return c.Update(rootContainerResourceId,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.SplitHorizontal(
			container.Top(container.PlaceWidget(v.clockSD)),
			container.Bottom(container.PlaceWidget(giveUpBtn)),
		),
	)
}

type MainView struct {
	db  *ent.Client
	svc *service.Service
}

func NewMainView(db *ent.Client, svc *service.Service) *MainView {
	return &MainView{db: db, svc: svc}
}

func (v MainView) Run() error {
	t, err := tcell.New()
	if err != nil {
		return err
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = v.db.Schema.Create(ctx, migrate.WithForeignKeys(false))
	if err != nil {
		return err
	}

	c, err := container.New(t, container.ID(rootContainerResourceId))
	if err != nil {
		return err
	}
	clock, err := NewClockView(v.svc, c)
	if err != nil {
		return err
	}
	clock.Run(ctx)
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	return termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(200*time.Millisecond))
}

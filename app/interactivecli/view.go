package interactivecli

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/go-gosh/tomato/app/ent"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

func NewClockView() (*ClockView, error) {
	clockSD, err := segmentdisplay.New()
	if err != nil {
		return nil, err
	}
	return &ClockView{
		clockSD: clockSD,
		tomato:  nil,
		alarmCh: make(chan struct{}),
	}, nil
}

type ClockView struct {
	clockSD *segmentdisplay.SegmentDisplay
	tomato  *ent.UserTomato
	alarmCh chan struct{}
}

func (v *ClockView) Tomato() *ent.UserTomato {
	return v.tomato
}

func (v *ClockView) SetTomato(tomato *ent.UserTomato) {
	v.tomato = tomato
}

func (v *ClockView) Run(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("exit")
				return
			case <-ticker.C:
				err := v.updated(ctx)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()
}

func (v *ClockView) updated(ctx context.Context) error {
	var chunk *segmentdisplay.TextChunk
	if v.tomato == nil {
		chunk = segmentdisplay.NewChunk(time.Now().Format("15:04:05"))
	} else {
		duration := v.tomato.RemainTime.Sub(time.Now()) / time.Second
		if duration <= 0 {
			v.alarmCh <- struct{}{}
			return nil
		}
		spacer := " "
		if duration%2 == 0 {
			spacer = ":"
		}
		chunk = segmentdisplay.NewChunk(
			fmt.Sprintf("%02d%s%02d", duration/60, spacer, duration%60),
		)
	}
	return v.clockSD.Write([]*segmentdisplay.TextChunk{chunk})
}

type MainView struct {
}

func (v MainView) Run() error {
	t, err := tcell.New()
	if err != nil {
		return err
	}
	defer t.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	clock, err := NewClockView()
	if err != nil {
		return err
	}
	clock.Run(ctx)
	go func() {
		<-time.After(1 * time.Second)
		clock.SetTomato(&ent.UserTomato{RemainTime: time.Now().Add(5 * time.Second)})
		<-clock.alarmCh
		for i := 0; i < 3; i++ {
			_ = exec.Command("afplay", "/System/Library/Sounds/Tink.aiff", "-v", "15", "-q", "1").Run()
		}
		cancel()
	}()
	c, err := container.New(
		t,
		container.Border(linestyle.Light),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.SplitHorizontal(
			container.Top(container.PlaceWidget(clock.clockSD)),
			container.Bottom(),
		),
	)
	if err != nil {
		return err
	}
	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	return termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(200*time.Millisecond))
}

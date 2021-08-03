package minicli

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
)

var (
	term    = termenv.ColorProfile()
	keyword = makeFgStyle("211")
	subtle  = makeFgStyle("241")
	dot     = colorFg(" • ", "236")
)

type frameMsg struct{}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

type Data struct {
	EndTime  time.Time
	Duration int
	Color    string
}

type DataManager interface {
	Alarm()

	Get() *Data
	Create() *Data
	GiveUp() *Data
	Submit(color string) *Data
	Close() *Data
}

type model struct {
	quit     bool
	cmd      tea.Cmd
	spinner  spinner.Model
	status   int
	data     *Data
	progress progress.Model
	dataMgr  DataManager
}

func NewModel(dataMgr DataManager) *model {
	s := spinner.NewModel()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	bar := progress.NewModel(progress.WithGradient("#B14FFF", "#00FFA3"))
	bar.Empty = bar.Full

	d := dataMgr.Get()
	m := &model{
		spinner:  s,
		data:     d,
		progress: bar,
		dataMgr:  dataMgr,
	}
	if d != nil {
		m.status = 1
	}
	return m
}

func (m model) Init() tea.Cmd {
	switch m.status {
	case 1:
		return frame()
	}
	return spinner.Tick
}

// Update Main update function.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.quit = true
			return m, tea.Quit
		}
	}
	switch m.status {
	case 0:
		// home page
		return m.updateHomePage(msg)
	case 1:
		// clock tick progress
		return m.updateClock(msg)
	case 2:
		// confirm give up
		return m.updateConfirmGiveUp(msg)
	case 3:
		// submit clock
		return m.updateSubmit(msg)
	}

	return m, m.cmd
}

// View The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.quit {
		return "\n  See you later!\n\n"
	}
	var tip string
	switch m.status {
	case 0:
		// home page
		s = m.viewHomePage()
	case 1:
		// clock tick progress
		s = m.viewClock()
		tip = subtle("p: pause") + dot
	case 2:
		// confirm give up
		s = m.viewConfirmGiveUp()
	case 3:
		// submit clock
		s = m.viewSubmit()
	}

	return indent.String("\n"+s+"\n\n"+tip+subtle("q, esc: quit")+"\n\n", 2)
}

func (m model) startTomatoView() string {
	diff := m.data.EndTime.Sub(time.Now()) / time.Second
	if diff <= 0 {
		diff = 0
	}
	return fmt.Sprintf("%s / %s", colorFg(fmt.Sprintf("%02d:%02d", diff/60, diff%60), "79"), fmt.Sprintf("%02d:%02d", m.data.Duration/60, m.data.Duration%60))
}

func (m model) updateHomePage(msg tea.Msg) (model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "s" {
		// start tomato
		m.data = m.dataMgr.Create()
		m.status = 1
		return m, frame()
	}
	m.spinner, m.cmd = m.spinner.Update(msg)
	return m, m.cmd
}

func (m model) viewHomePage() string {
	return "Ready to make a tomato clock?\n\nPress " +
		keyword("s") +
		" to start  " + m.spinner.View()
}

func (m model) updateClock(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.data.EndTime.Before(time.Now()) {
		// alarm clock
		m.dataMgr.Alarm()
		// switch submit view
		m.status = 3
		return m, spinner.Tick
	}
	if msg, ok := msg.(tea.KeyMsg); ok {
		if msg.String() == "p" {
			// switch give up tomato view
			m.status = 2
			return m, m.cmd
		}
	}
	return m, frame()
}

func (m model) viewConfirmGiveUp() string {
	return "Give up this tomato clock?\n\nPress " +
		colorFg("Y for yes, N or any for no    ", "79") + m.spinner.View()
}

func (m model) updateConfirmGiveUp(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "y" {
			// give up tomato
			m.status = 0
			m.data = m.dataMgr.GiveUp()
			m.spinner, m.cmd = m.spinner.Update(msg)
			return m, m.cmd
		}
		m.status = 1
		return m, frame()
	}
	m.spinner, m.cmd = m.spinner.Update(msg)
	return m, m.cmd
}

func (m model) viewClock() string {
	var per float64
	diff := m.data.EndTime.Sub(time.Now()) / time.Second
	if diff <= 0 {
		per = 0
	} else {
		per = float64(diff) / float64(m.data.Duration)
	}

	return m.startTomatoView() + "\n\n" + m.progress.ViewAs(per)
}

func (m model) updateSubmit(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "y" {
			// submit tomato and start a new
			m.data = m.dataMgr.Submit(m.data.Color)
			// switch to clock
			m.status = 1
			return m, frame()
		}
		if key.String() == "n" {
			// submit tomato only
			m.data = m.dataMgr.Close()
			// switch to home page
			m.status = 0
		}
	}
	m.spinner, m.cmd = m.spinner.Update(msg)
	return m, m.cmd
}

func (m model) viewSubmit() string {
	return "You are completed a tomato, for more?\n\nPress " +
		colorFg("Y for yes, N for no    ", "79") + m.spinner.View()
}

// Color a string's foreground with the given value.
func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}

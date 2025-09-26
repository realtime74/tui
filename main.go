package main

import (
	"fmt"
	"os"
	"time"

	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
	"github.com/realtime74/tui/ui"
)

type Component interface {
	Render()
}

var _model = make(map[string]Component)

func _down() {
	bar := _model["rbar"].(*ui.VBar)
	bar.Down()
}

func _up() {
	bar := _model["rbar"].(*ui.VBar)
	bar.Up()
}

var _messages = make(chan tcell.Event, 10)

func _poll(s1 tcell.Screen) tcell.Event {
	select {
	case ev := <-_messages:
		return ev
	}
}

func _loop(s1 tcell.Screen) bool {
	s1.Show()
	ev := _poll(s1)
	switch ev := ev.(type) {
	case TimerEvent:
		scr.DrawText(s1, 0, 0,
			fmt.Sprintf("Tick: %d", ev.Cycle),
			tcell.StyleDefault)
		ball := _model["ball"].(*ui.Ball)
		ball.Move()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape ||
			ev.Key() == tcell.KeyCtrlC {
			return false
		}
		switch ev.Rune() {
		case 'j', 'J':
			_down()
		case 'k', 'K':
			_up()
		case 'q', 'Q':
			return false
		}
		return true
	case *tcell.EventResize:
		s1.Sync()
	}
	return true
}

var logger *slog.Logger

func main() {

	file, err := os.OpenFile(
		"logfile.json",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger = slog.New(slog.NewJSONHandler(file, nil))
	defer file.Close()

	s1 := scr.New()
	defer s1.Fini()

	s1.Clear()

	greeting := "Hello, World!"
	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlue)

	w, h := s1.Size()
	scr.DrawText(s1, w/2-len(greeting)/2, h/2, greeting, style)

	header := ui.NewHBar(s1, 0, 0, w)
	header.SetText("My TUI App")
	status := ui.NewHBar(s1, 0, h-1, w)
	status.SetText("ESC/q: quit  j/k: up/down")

	height := h/5 - 1
	xpos := h/2 - height/2

	lbar := ui.NewVBar(s1, 0, xpos, height)
	lbar.Render()

	rbar := ui.NewVBar(s1, w-1, xpos, height)
	rbar.Render()

	playground := ui.NewPlayground(s1, 1, 1, w-2, h-2)
	playground.Render()

	ball := ui.NewBall(s1, playground, w/2, h/2)
	ball.Render()

	_model["playground"] = playground
	_model["ball"] = ball

	_model["lbar"] = lbar
	_model["rbar"] = rbar

	go _timerThread()

	go func() {
		for {
			msg := s1.PollEvent()
			_messages <- msg
		}
	}()

	for {
		if !_loop(s1) {
			break
		}
	}
}

type TimerEvent struct {
	Cycle int
}

func _timerThread() {
	logger.Info("Starting timer thread")
	tick := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		tick++
		select {
		case <-ticker.C:
			_messages <- TimerEvent{Cycle: tick}
		}
	}
}

func (e TimerEvent) When() time.Time {
	return time.Now()
}

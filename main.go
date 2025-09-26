package main

import (
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

func _loop(s1 tcell.Screen) bool {
	s1.Show()
	ev := s1.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
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

func main() {
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

	_model["lbar"] = lbar
	_model["rbar"] = rbar

	for {
		if !_loop(s1) {
			break
		}
	}
}

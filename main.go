package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic("tcell setup")
	}
	if err := s.Init(); err != nil {
		panic("tcell init")
	}
	defer s.Fini()
	s.Clear()

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)
	greeting := "Hello, World!"

	w, h := s.Size()
	scr.DrawText(s, w/2-len(greeting)/2, h/2, greeting, style)

	s.Show()
	time.Sleep(5 * time.Second)
}

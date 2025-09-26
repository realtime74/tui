package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
	"github.com/realtime74/tui/ui"
)

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

	s1.Show()
	time.Sleep(5 * time.Second)
}

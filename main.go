package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawText(s tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}

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
	drawText(s, w/2-len(greeting)/2, h/2, greeting, style)

	s.Show()
	time.Sleep(5 * time.Second)
}

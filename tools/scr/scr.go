package scr

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

var InitError error

func New() tcell.Screen {

	s, err := tcell.NewScreen()
	if err != nil {
		InitError = fmt.Errorf("failed to initialize screen: %v", err)
		panic(InitError)
	}
	if err := s.Init(); err != nil {
		InitError = fmt.Errorf("failed to initialize screen: %v", err)
		panic(InitError)
	}
	return s
}

func Fill(s tcell.Screen, x, y, w int, ch rune, style tcell.Style) {
	for ii := 0; ii < w; ii++ {
		s.SetContent(x+ii, y, ch, nil, style)
	}
}

func DrawText(s tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}

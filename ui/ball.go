package ui

import (
	"github.com/gdamore/tcell/v2"
)

type Ball struct {
	X, Y int

	scr        tcell.Screen
	Playground *Playground
}

func NewBall(scr tcell.Screen, g *Playground, x, y int) *Ball {
	return &Ball{
		scr:        scr,
		X:          x,
		Y:          y,
		Playground: g,
	}
}

func (b *Ball) Render() {
	style := tcell.StyleDefault.
		Foreground(tcell.NewRGBColor(63, 63, 70)).
		Background(tcell.ColorBlack)

	b.scr.SetContent(b.X, b.Y, '⏺', nil, style)
}

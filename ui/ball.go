package ui

import (
	"github.com/gdamore/tcell/v2"
)

type Ball struct {
	X, Y      int
	direction int

	scr        tcell.Screen
	Playground *Playground
}

func NewBall(scr tcell.Screen, g *Playground, x, y int) *Ball {
	return &Ball{
		direction:  1,
		scr:        scr,
		X:          x,
		Y:          y,
		Playground: g,
	}
}

func (b *Ball) Move() {
	b.scr.SetContent(b.X, b.Y, ' ', nil,
		tcell.StyleDefault.Background(tcell.ColorBlack).
			Foreground(tcell.ColorWhite))
	b.X += b.direction
	rwall := b.Playground.X + b.Playground.W - 1
	if b.X >= rwall {
		b.X -= b.direction
		b.direction = -b.direction
	}
	if b.X <= b.Playground.X {
		b.X -= b.direction
		b.direction = -b.direction
	}
	b.Render()
}

func (b *Ball) Render() {
	style := tcell.StyleDefault.
		Foreground(tcell.NewRGBColor(63, 63, 70)).
		Background(tcell.ColorBlack)

	b.scr.SetContent(b.X, b.Y, '⏺', nil, style)
}

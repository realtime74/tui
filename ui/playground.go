package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
)

type Playground struct {
	scr        tcell.Screen
	X, Y, W, H int
}

func NewPlayground(scr tcell.Screen, x, y, w, h int) *Playground {
	return &Playground{
		scr: scr,
		X:   x,
		Y:   y,
		W:   w,
		H:   h,
	}
}

func (p *Playground) Render() {
	style := tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack)

	for ii := 0; ii < p.H; ii++ {
		scr.Fill(p.scr, p.X, p.Y+ii, p.W, ' ', style)
	}
}

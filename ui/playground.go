package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
)

type Playground struct {
	scr        tcell.Screen
	x, y, w, h int
}

func NewPlayground(scr tcell.Screen, x, y, w, h int) *Playground {
	return &Playground{
		scr: scr,
		x:   x,
		y:   y,
		w:   w,
		h:   h,
	}
}

func (p *Playground) Render() {
	style := tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorBlack)

	for ii := 0; ii < p.h; ii++ {
		scr.Fill(p.scr, p.x, p.y+ii, p.w, ' ', style)
	}
}

package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
)

type VBar struct {
	X, Y, H int

	scr tcell.Screen
}

func NewVBar(scr tcell.Screen, x, y, h int) *VBar {
	return &VBar{
		X:   x,
		Y:   y,
		H:   h,
		scr: scr,
	}
}

func (v *VBar) Render() {
	style := tcell.StyleDefault.
		Background(tcell.ColorWhite)
	scr.HFill(v.scr, v.X, v.Y, v.H, ' ', style)
}

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

func (v *VBar) Down() {
	_, height := v.scr.Size()
	if v.Y+v.H < height-1 {
		v.Y++
	}
	v.Render()
}

func (v *VBar) Up() {
	if v.Y > 2 {
		v.Y--
	}
	v.Render()
}

func (v *VBar) Render() {
	bgstyle := tcell.StyleDefault.
		Background(tcell.ColorBlack)
	style := tcell.StyleDefault.
		Background(tcell.ColorWhite)

	_, height := v.scr.Size()

	scr.HFill(v.scr, v.X, 1, v.Y-1, ' ', bgstyle)
	scr.HFill(v.scr, v.X, v.Y, v.H, ' ', style)
	scr.HFill(v.scr, v.X, v.Y+v.H, (height-1)-(v.Y+v.H), ' ', bgstyle)
}

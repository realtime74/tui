package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/tui/tools/scr"
)

type VBar struct {
	X, Y, W int
	Text string

	scr  tcell.Screen
}

func NewVBar(scr tcell.Screen, x, y, w int) *VBar {
	return &VBar{
		X:    x,
		Y:    y,
		W:    w,
		Text: "",
		scr:  scr,
	}
}

func (v *VBar) SetText(text string) {
	v.Text = text
	v._render()
}

func (v *VBar) _render() {
	style := tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlue)

	scr.DrawText(v.scr, v.X, v.Y, v.Text, style)
	scr.Fill(v.scr, v.X+len(v.Text), v.Y, v.W-len(v.Text), ' ', style)
}

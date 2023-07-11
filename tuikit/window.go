package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Container
	hasBorder bool
}

func (w *Window) Draw() {
	if w.GetScreen() == nil {
		panic("Select's Screen is not defined")
	}
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	y := w.y
	for r := 0; r < w.height; r++ {
		x := w.x
		for c := 0; c < w.width; c++ {
			if !w.hasBorder {
				continue
			}
			char := ' ' // default character is a space
			if r == 0 { // top border
				if c == 0 {
					char = '┌'
				} else if c == w.width-1 {
					char = '┐'
				} else {
					char = '─'
				}
			} else if r == w.height-1 { // bottom border
				if c == 0 {
					char = '└'
				} else if c == w.width-1 {
					char = '┘'
				} else {
					char = '─'
				}
			} else if c == 0 { // left border
				char = '│'
			} else if c == w.width-1 { // right border
				char = '│'
			}
			w.GetScreen().Screen.SetContent(x+c, y+r, char, nil, borderStyle)
		}
	}
}

func NewWindow() *Window {
	return &Window{
		Container: *NewContainer(),
	}
}

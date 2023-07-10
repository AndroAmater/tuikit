package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type Window struct {
	Container
	elements []Element
	width    int
	height   int
	x        int
	y        int
	growX    bool
	growY    bool
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

func (w *Window) SetScreen(screen *Screen) {
	w.SetScreen(screen)
}

func (w *Window) HandleEvent(event tcell.Event) {
	for _, element := range w.elements {
		element.HandleEvent(event)
	}
}

func (w *Window) AddElement(element Element) {
	w.elements = append(w.elements, element)
}

func NewWindow(x int, y int, width int, height int, growX bool, growY bool) *Window {
	return &Window{
		Container: *NewContainer(),
		x:         x,
		y:         y,
		width:     width,
		height:    height,
		growX:     growX,
		growY:     growY,
	}
}

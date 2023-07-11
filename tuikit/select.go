package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type Select struct {
	Element
	choices    []string
	selected   []bool
	indicators [4]string
	cursor     int
}

func (s *Select) Draw() {
	if s.screen == nil {
		panic("Select's Screen is not defined")
	}
	y := s.y
	for i, choice := range s.choices {
		indicator := s.indicators[0]
		if s.selected[i] {
			indicator = s.indicators[1]
		}
		line := s.indicators[3] + indicator + choice
		if i == s.cursor {
			line = s.indicators[2] + indicator + choice
		}

		x := s.x
		for _, r := range line {
			s.screen.Screen.SetContent(x, y, r, nil, tcell.StyleDefault)
			x++
		}
		y++
	}
}

func (s *Select) HandleEvent(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			if s.cursor > 0 {
				s.cursor--
			}
		case tcell.KeyDown:
			if s.cursor < len(s.choices)-1 {
				s.cursor++
			}
		case tcell.KeyRune:
			if ev.Rune() == ' ' {
				s.selected[s.cursor] = !s.selected[s.cursor]
			}
		}
	}
	s.eventHandler(event)
}

func NewSelect(choices []string) *Select {
	return &Select{
		Element:    *NewElement(),
		choices:    choices,
		selected:   make([]bool, len(choices)),
		indicators: [4]string{"[ ] ", "[*] ", "> ", "  "},
		cursor:     0,
	}
}

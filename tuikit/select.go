package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type Select struct {
	Element
	choices      []string
	selected     []bool
	width        int
	height       int
	indicators   [4]string
	screen       *Screen
	y            int
	x            int
	growX        bool
	growY        bool
	eventHandler func(event tcell.Event)
	cursor       int
}

func NewSelect(choices []string, selected []bool) *Select {
	return &Select{
		choices:      choices,
		selected:     selected,
		width:        0,
		height:       len(choices),
		indicators:   [4]string{"[ ] ", "[*] ", "> ", "  "},
		screen:       nil,
		y:            0,
		x:            0,
		growX:        true,
		growY:        false,
		eventHandler: func(event tcell.Event) {},
		cursor:       0,
	}
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

func (s *Select) setScreen(screen *Screen) {
	s.screen = screen
	if s.growX {
		width, _ := screen.Screen.Size()
		s.width = width
	}
	if s.growY {
		_, height := screen.Screen.Size()
		s.height = height
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

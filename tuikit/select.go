package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type SelectIndicators struct {
	checkbox         string
	checkboxSelected string
	cursor           string
	cursorSelected   string
}

type Select struct {
	Element
	choices    []string
	selected   []bool
	indicators SelectIndicators
	cursor     int
	elements   []SelectOption
}

type SelectOption struct {
	Element
	selected   bool
	indicators *SelectIndicators
	focused    bool
	choice     string
}

func newSelectOption(
	indicators *SelectIndicators,
	selected bool,
	focused bool,
	choice string,
) *SelectOption {
	return &SelectOption{
		Element:    *NewElement(),
		indicators: indicators,
		selected:   selected,
		focused:    focused,
		choice:     choice,
	}
}

func (s *SelectOption) Draw() {
	content := s.choice
	if s.selected {
		content = s.indicators.checkboxSelected + content
	} else {
		content = s.indicators.checkbox + content
	}
	if s.focused {
		content = s.indicators.cursorSelected + content
	} else {
		content = s.indicators.cursor + content
	}
	s.SetContent(content)
	s.Element.Draw()
}

func (s *SelectOption) SetSelected(selected bool) {
	s.selected = selected
}

func (s *SelectOption) SetFocused(focused bool) {
	s.focused = focused
}

func (s *Select) HandleEvent(event tcell.Event) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			if s.cursor > 0 {
				s.SetFocused(s.cursor - 1)
			}
			break
		case tcell.KeyDown:
			if s.cursor < len(s.choices)-1 {
				s.SetFocused(s.cursor + 1)
			}
			break
		case tcell.KeyRune:
			if ev.Rune() == ' ' {
				s.ToggleSelected(s.cursor)
			} else if ev.Rune() == 'k' {
				if s.cursor > 0 {
					s.SetFocused(s.cursor - 1)
				}
			} else if ev.Rune() == 'j' {
				if s.cursor < len(s.choices)-1 {
					s.SetFocused(s.cursor + 1)
				}
			}
			break
		}
	}
	s.eventHandler(event)
}

func (s *Select) ToggleSelected(i int) {
	s.selected[i] = !s.selected[i]
	so, ok := s.GetChildren()[i].(*SelectOption)
	if !ok {
		panic("Error: Select child is not a SelectOption")
	}
	so.SetSelected(s.selected[i])
	s.Draw()
}

func (s *Select) SetFocused(i int) {
	so, ok := s.GetChildren()[s.cursor].(*SelectOption)
	if !ok {
		panic("Error: Select child is not a SelectOption")
	}
	so.SetFocused(false)
	s.cursor = i
	so, ok = s.GetChildren()[s.cursor].(*SelectOption)
	if !ok {
		panic("Error: Select child is not a SelectOption")
	}
	so.SetFocused(true)
	s.Draw()
}

func (s *Select) SetChoices(choices []string, selected []bool) {
	s.RemoveAllChildren()
	s.selected = selected
	for i, choice := range choices {
		so := newSelectOption(&s.indicators, selected[i], i == 0, choice)
		so.SetHeight(1)
		so.SetGrowX(true)
		s.AddChild(so)
	}
	s.choices = choices
	s.selected = selected
}

func NewSelect(choices []string, selected []bool) *Select {
	if len(selected) < len(choices) {
		for i := len(selected); i < len(choices); i++ {
			selected = append(selected, false)
		}
	}
	s := &Select{
		Element:  *NewElement(),
		choices:  choices,
		selected: make([]bool, len(choices)),
		indicators: SelectIndicators{
			checkbox:         "[ ] ",
			checkboxSelected: "[*] ",
			cursor:           "  ",
			cursorSelected:   "> ",
		},
		cursor: 0,
	}
	s.SetContentDirection(ContainerDirection.Column)
	s.SetChoices(choices, selected)
	return s
}

var (
	_ IsElement = &Select{}
	_ IsElement = &SelectOption{}
)

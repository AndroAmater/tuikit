package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	HasScreen
	ContainsElements
	Screen      tcell.Screen
	ExitChannel *chan struct{}
	elements    []IsElement
	screen      *Screen
}

func NewScreen() (*Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	s := &Screen{
		Screen: screen,
	}

	s.screen = s

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.Screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					close(quit)
					return
				}
				for _, e := range s.elements {
					e.HandleEvent(ev)
				}
			}
			s.Draw()
		}
	}()

	s.ExitChannel = &quit

	return s, err
}

func setScreen(s *Screen) {
	s.screen = s
}

func (s *Screen) AddChild(e IsElement) {
	s.elements = append(s.elements, e)
	e.setScreen(s)
}

func (s *Screen) GetChildren() []IsElement {
	return s.elements
}

func (s *Screen) DrawChildren() {
	for _, e := range s.elements {
		e.Draw()
	}
}

func (s *Screen) Draw() {
	s.Screen.Clear()
	s.DrawChildren()
	s.Screen.Show()
}

func (s *Screen) Close() {
	s.Screen.Fini()
}

func (s *Screen) GetHeight() int {
	_, height := s.Screen.Size()
	// FIXME: Why is size always 1 less than the actual size?
	return height - 1
}

func (s *Screen) GetWidth() int {
	width, _ := s.Screen.Size()
	// FIXME: Why is size always 1 less than the actual size?
	return width - 1
}

func (s *Screen) GetSize() (int, int) {
	return s.GetHeight(), s.GetWidth()
}

package tuikit

import "github.com/gdamore/tcell/v2"

type Element struct {
	HasScreen
	HandlesEvents
	container    ContainsElements
	eventHandler func(event tcell.Event)
	screen       *Screen
}

func (e *Element) Draw() {}

func (e *Element) setContainer(c ContainsElements) {
	e.container = c
}

func (e *Element) GetContainer() ContainsElements {
	return e.container
}

func (e *Element) HandleEvent(event tcell.Event) {
	e.GetEventHandler()(event)
}

func (e *Element) SetEventHandler(handler func(event tcell.Event)) {
	e.eventHandler = handler
}

func (e *Element) GetEventHandler() func(event tcell.Event) {
	return e.eventHandler
}

func (e *Element) setScreen(s *Screen) {
	e.screen = s
}

func (e *Element) GetScreen() *Screen {
	return e.screen
}

func NewElement() *Element {
	return &Element{
		container: nil,
	}
}

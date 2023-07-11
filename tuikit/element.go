package tuikit

import "github.com/gdamore/tcell/v2"

type Element struct {
	container    ContainsElements
	eventHandler func(event tcell.Event)
	screen       *Screen
	y            int
	x            int
	growX        bool
	growY        bool
	width        int
	height       int
}

// IsElement interface

func (e *Element) Draw() {}

// HandlesEvents interface

func (e *Element) HandleEvent(event tcell.Event) {
	e.GetEventHandler()(event)
}

func (e *Element) SetEventHandler(handler func(event tcell.Event)) {
	e.eventHandler = handler
}

func (e *Element) GetEventHandler() func(event tcell.Event) {
	return e.eventHandler
}

// HasContainer interface

func (e *Element) setContainer(c ContainsElements) {
	e.container = c
}

func (e *Element) GetContainer() ContainsElements {
	return e.container
}

// HasPosition interface

func (e *Element) GetX() int {
	return e.x
}

func (e *Element) SetX(x int) {
	e.x = x
}

func (e *Element) GetY() int {
	return e.y
}

func (e *Element) SetY(y int) {
	e.y = y
}

func (e *Element) GetPosition() (int, int) {
	return e.GetX(), e.GetY()
}

func (e *Element) SetPosition(x int, y int) {
	e.SetX(x)
	e.SetY(y)
}

// HasScreen interface

func (e *Element) setScreen(s *Screen) {
	e.screen = s
}

func (e *Element) GetScreen() *Screen {
	return e.screen
}

// HasSize interface

func (e *Element) GetWidth() int {
	return e.width
}

func (e *Element) SetWidth(width int) {
	e.width = width
}

func (e *Element) GetHeight() int {
	return e.height
}

func (e *Element) SetHeight(height int) {
	e.height = height
}

func (e *Element) GetSize() (int, int) {
	return e.GetWidth(), e.GetHeight()
}

func (e *Element) SetSize(width int, height int) {
	e.SetWidth(width)
	e.SetHeight(height)
}

func (e *Element) GetGrowX() bool {
	return e.growX
}

func (e *Element) SetGrowX(growX bool) {
	e.growX = growX
}

func (e *Element) GetGrowY() bool {
	return e.growY
}

func (e *Element) SetGrowY(growY bool) {
	e.growY = growY
}

// Other methods

func NewElement() *Element {
	return &Element{
		container: nil,
	}
}

// Validate interface implementation

var _ IsElement = &Element{}

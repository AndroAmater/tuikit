package tuikit

import "github.com/gdamore/tcell/v2"

type Container struct {
	Element
	elements []IsElement
}

// ContainsElements interface

func (c *Container) AddElement(e IsElement) {
	c.elements = append(c.elements, e)
	e.setContainer(c)
	e.setScreen(c.GetScreen())
}

func (c *Container) GetElements() []IsElement {
	return c.elements
}

func (c *Container) DrawChildren() {
	for _, e := range c.elements {
		e.Draw()
	}
}

// HandlesEvents interface

func (c *Container) HandleEvent(event tcell.Event) {
	for _, e := range c.elements {
		e.HandleEvent(event)
	}
	c.eventHandler(event)
}

// Other methods

func NewContainer() *Container {
	return &Container{
		elements: make([]IsElement, 0),
	}
}

// Validate interface implementations

var (
	_ ContainsElements = &Container{}
	_ IsElement        = &Container{}
)

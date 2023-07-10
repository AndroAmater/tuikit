package tuikit

type Container struct {
	Element
	ContainsElements
	elements []IsElement
}

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

func NewContainer() *Container {
	return &Container{
		elements: make([]IsElement, 0),
	}
}

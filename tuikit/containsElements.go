package tuikit

type ContainsElements interface {
	// Adds an element to the container
	AddElement(e IsElement)
	// Removes an element from the container
	GetElements() []IsElement
	// Draws the container's children
	DrawChildren()
}

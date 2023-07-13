package tuikit

type IsElement interface {
	HandlesEvents
	HasContainer
	HasPosition
	HasScreen
	HasSize
	ContainsElements
	// Draws the element to the screen
	Draw()
}

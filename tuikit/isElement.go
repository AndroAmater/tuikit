package tuikit

type IsElement interface {
	HandlesEvents
	HasContainer
	HasPosition
	HasScreen
	HasSize
	// Draws the element to the screen
	Draw()
}

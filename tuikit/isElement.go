package tuikit

/*
 * TODO:
 * - Add overflow
 * - Add position
 * - Add content (text)
 * - Add scroll
 * - Add wrap
 */

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

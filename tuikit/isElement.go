package tuikit

import "github.com/google/uuid"

/*
 * TODO:
 * - Add overflow
 * - Add position
 * - Add content (text)
 * - Add scroll
 * - Add wrap
 * - Add removing children
 */

type IsElement interface {
	HandlesEvents
	HasContainer
	HasPosition
	HasScreen
	HasSize
	ContainsElements
	// Draws the element to the screen
	HasContent
	GetUUID() uuid.UUID
	Draw()
}

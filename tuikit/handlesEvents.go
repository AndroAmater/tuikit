package tuikit

import "github.com/gdamore/tcell/v2"

type HandlesEvents interface {
	// Handles the provided event with the element event handlers
	HandleEvent(event tcell.Event)
	// Sets the user event handler for the element
	SetEventHandler(handler func(event tcell.Event))
	// Gets the user event handler for the element
	GetEventHandler() func(event tcell.Event)
}

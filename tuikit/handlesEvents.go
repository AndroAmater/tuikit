package tuikit

import "github.com/gdamore/tcell/v2"

type HandlesEvents interface {
	HandleEvent(event tcell.Event)
	SetEventHandler(handler func(event tcell.Event))
	GetEventHandler() func(event tcell.Event)
}

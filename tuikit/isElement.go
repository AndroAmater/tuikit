package tuikit

type IsElement interface {
	HandlesEvents
	HasScreen
	HasContainer
	HasSize
	HasPosition
	Draw()
}

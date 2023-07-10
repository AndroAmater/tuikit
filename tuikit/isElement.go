package tuikit

type IsElement interface {
	HandlesEvents
	HasScreen
	HasContainer
	Draw()
}

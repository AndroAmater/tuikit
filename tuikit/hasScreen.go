package tuikit

type HasScreen interface {
	setScreen(s *Screen)
	GetScreen() *Screen
}

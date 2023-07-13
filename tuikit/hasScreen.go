package tuikit

type HasScreen interface {
	// Sets the screen the element is attached to
	setScreen(s *Screen)
	// Gets the screen the element is attached to
	GetScreen() *Screen
}

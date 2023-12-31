package tuikit

type HasPosition interface {
	// Gets the element's X relative position
	GetX() int
	// Gets the element's Y relative position
	GetY() int
	// Gets the element's X absolute position
	GetAbsoluteX() int
	// Gets the element's Y absolute position
	GetAbsoluteY() int
	// Gets the element's X and Y relative positions
	GetPosition() (int, int)
	// Gets the element's X and Y absolute positions
	GetAbsolutePosition() (int, int)
	// Sets the element's X relative position
	SetX(x int)
	// Sets the element's Y relative position
	SetY(y int)
	// Sets the element's X and Y relative positions
	SetPosition(x, y int)
}

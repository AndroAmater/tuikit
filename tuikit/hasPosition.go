package tuikit

type HasPosition interface {
	GetX() int
	GetY() int
	GetPosition() (int, int)
	SetX(x int)
	SetY(y int)
	SetPosition(x, y int)
}

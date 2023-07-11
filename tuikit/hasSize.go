package tuikit

type HasSize interface {
	GetWidth() int
	GetHeight() int
	GetSize() (int, int)
	SetWidth(int)
	SetHeight(int)
	SetSize(int, int)
}

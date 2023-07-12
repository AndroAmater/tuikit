package tuikit

type BorderStyleType int

const (
	none BorderStyleType = iota
	single
)

var BorderStyle = struct {
	None   BorderStyleType
	Single BorderStyleType
}{
	None:   none,
	Single: single,
}

type BorderOpts struct {
	Style BorderStyleType
	Color string
}

type HasSize interface {
	GetInnerWidth() int
	GetInnerHeight() int
	GetInnerSize() (int, int)
	GetWidth() int
	GetHeight() int
	GetSize() (int, int)
	GetOuterWidth() int
	GetOuterHeight() int
	GetOuterSize() (int, int)
	SetWidth(int)
	SetHeight(int)
	SetSize(int, int)

	GetGrowX() bool
	GetGrowY() bool
	GetGrow() (bool, bool)
	SetGrowX(bool)
	SetGrowY(bool)
	SetGrow(bool, bool)

	GetMarginTop() int
	GetMarginBottom() int
	GetMarginLeft() int
	GetMarginRight() int
	GetMargin() (int, int, int, int)
	SetMarginTop(int)
	SetMarginBottom(int)
	SetMarginLeft(int)
	SetMarginRight(int)
	SetMargin(int, int, int, int)

	GetPaddingTop() int
	GetPaddingBottom() int
	GetPaddingLeft() int
	GetPaddingRight() int
	GetPadding() (int, int, int, int)
	SetPaddingTop(int)
	SetPaddingBottom(int)
	SetPaddingLeft(int)
	SetPaddingRight(int)
	SetPadding(int, int, int, int)

	GetBorderTop() BorderOpts
	GetBorderBottom() BorderOpts
	GetBorderLeft() BorderOpts
	GetBorderRight() BorderOpts
	GetBorder() (BorderOpts, BorderOpts, BorderOpts, BorderOpts)
	SetBorderTop(BorderStyleType, string)
	SetBorderBottom(BorderStyleType, string)
	SetBorderLeft(BorderStyleType, string)
	SetBorderRight(BorderStyleType, string)
	SetBorder(BorderOpts, BorderOpts, BorderOpts, BorderOpts)
}

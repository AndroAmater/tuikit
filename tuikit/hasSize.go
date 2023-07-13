package tuikit

type BorderStyleType int

const (
	none BorderStyleType = iota
	single
)

// Border styles enum
var BorderStyle = struct {
	None   BorderStyleType
	Single BorderStyleType
}{
	None:   none,
	Single: single,
}

// Border configuration options
type BorderOpts struct {
	Style BorderStyleType
	Color string
}

type HasSize interface {
	// Getts the element's width without margins, borders and padding
	GetInnerWidth() int
	// Gets the element's height without margins, borders and padding
	GetInnerHeight() int
	// Gets the element's width and height without margins, borders and padding
	GetInnerSize() (int, int)
	// Gets the element's width including borders and padding
	GetWidth() int
	// Gets the element's height including borders and padding
	GetHeight() int
	// Gets the element's width and height including borders and padding
	GetSize() (int, int)
	// Gets the element's width including margins, borders and padding
	GetOuterWidth() int
	// Gets the element's height including margins, borders and padding
	GetOuterHeight() int
	// Gets the element's width and height including margins, borders and padding
	GetOuterSize() (int, int)
	// Sets the element's width (without margins, with borders and paddding,
	// content size will be calculated automatically)
	SetWidth(int)
	// Sets the element's height (without margins, with borders and paddding,
	// content size will be calculated automatically)
	SetHeight(int)
	// Sets the element's width and height (without margins, with borders and
	// paddding, content size will be calculated automatically)
	SetSize(int, int)

	// Gets the element's GrowX flag (will the element grow to fill container horizontally)
	GetGrowX() bool
	// Gets the element's GrowY flag (will the element grow to fill container vertically)
	GetGrowY() bool
	// Gets the element's GrowX and GrowY flags (will the element grow to fill container)
	GetGrow() (bool, bool)
	// Sets the element's GrowX flag (will the element grow to fill container horizontally)
	SetGrowX(bool)
	// Sets the element's GrowY flag (will the element grow to fill container vertically)
	SetGrowY(bool)
	// Sets the element's GrowX and GrowY flags (will the element grow to fill container)
	SetGrow(bool, bool)

	// Gets the element's top margin size
	GetMarginTop() int
	// Gets the element's bottom margin size
	GetMarginBottom() int
	// Gets the element's left margin size
	GetMarginLeft() int
	// Gets the element's right margin size
	GetMarginRight() int
	// Gets the element's margin sizes (top, right, bottom, left)
	GetMargin() (int, int, int, int)
	// Sets the element's top margin size
	SetMarginTop(int)
	// Sets the element's bottom margin size
	SetMarginBottom(int)
	// Sets the element's left margin size
	SetMarginLeft(int)
	// Sets the element's right margin size
	SetMarginRight(int)
	// Sets the element's margin sizes (top, right, bottom, left)
	SetMargin(int, int, int, int)

	// Gets the element's top padding size
	GetPaddingTop() int
	// Gets the element's bottom padding size
	GetPaddingBottom() int
	// Gets the element's left padding size
	GetPaddingLeft() int
	// Gets the element's right padding size
	GetPaddingRight() int
	// Gets the element's padding sizes (top, right, bottom, left)
	GetPadding() (int, int, int, int)
	// Sets the element's top padding size
	SetPaddingTop(int)
	// Sets the element's bottom padding size
	SetPaddingBottom(int)
	// Sets the element's left padding size
	SetPaddingLeft(int)
	// Sets the element's right padding size
	SetPaddingRight(int)
	// Sets the element's padding sizes (top, right, bottom, left)
	SetPadding(int, int, int, int)

	// Gets the element's top border style and color
	GetBorderTop() BorderOpts
	// Gets the element's bottom border style and color
	GetBorderBottom() BorderOpts
	// Gets the element's left border style and color
	GetBorderLeft() BorderOpts
	// Gets the element's right border style and color
	GetBorderRight() BorderOpts
	// Gets the element's border styles and colors (top, right, bottom, left)
	GetBorder() (BorderOpts, BorderOpts, BorderOpts, BorderOpts)
	// Sets the element's top border style and color
	SetBorderTop(BorderStyleType, string)
	// Sets the element's bottom border style and color
	SetBorderBottom(BorderStyleType, string)
	// Sets the element's left border style and color
	SetBorderLeft(BorderStyleType, string)
	// Sets the element's right border style and color
	SetBorderRight(BorderStyleType, string)
	// Sets the element's border styles and colors (top, right, bottom, left)
	SetBorder(BorderOpts, BorderOpts, BorderOpts, BorderOpts)
}

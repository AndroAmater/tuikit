package tuikit

import (
	"github.com/gdamore/tcell/v2"
)

type PixelTypeType int

const (
	margin PixelTypeType = iota
	padding
	topLeftBorder
	topBorder
	topRightBorder
	leftBorder
	rightBorder
	bottomLeftBorder
	bottomBorder
	bottomRightBorder
	content
)

var PixelType = struct {
	Margin            PixelTypeType
	Padding           PixelTypeType
	TopLeftBorder     PixelTypeType
	TopBorder         PixelTypeType
	TopRightBorder    PixelTypeType
	LeftBorder        PixelTypeType
	RightBorder       PixelTypeType
	BottomLeftBorder  PixelTypeType
	BottomBorder      PixelTypeType
	BottomRightBorder PixelTypeType
	Content           PixelTypeType
}{
	Margin:            margin,
	Padding:           padding,
	TopLeftBorder:     topLeftBorder,
	TopBorder:         topBorder,
	TopRightBorder:    topRightBorder,
	LeftBorder:        leftBorder,
	RightBorder:       rightBorder,
	BottomLeftBorder:  bottomLeftBorder,
	BottomBorder:      bottomBorder,
	BottomRightBorder: bottomRightBorder,
	Content:           content,
}

type Window struct {
	Container
	hasBorder bool
}

var PixelRuneMap = map[PixelTypeType]rune{
	PixelType.Margin:            ' ',
	PixelType.Padding:           ' ',
	PixelType.TopLeftBorder:     '┌',
	PixelType.TopBorder:         '─',
	PixelType.TopRightBorder:    '┐',
	PixelType.LeftBorder:        '│',
	PixelType.RightBorder:       '│',
	PixelType.BottomLeftBorder:  '└',
	PixelType.BottomBorder:      '─',
	PixelType.BottomRightBorder: '┘',
	PixelType.Content:           ' ',
}

func getPositionPixelType(w *Window, x int, y int) PixelTypeType {
	// Margin
	if y <= w.GetMarginTop() {
		return PixelType.Margin
	}
	if y > w.GetMarginTop()+w.GetHeight() {
		return PixelType.Margin
	}
	if x <= w.GetMarginLeft() {
		return PixelType.Margin
	}
	if x > w.GetMarginLeft()+w.GetWidth() {
		return PixelType.Margin
	}
	// Top border
	if y == w.GetMarginTop()+1 && w.GetBorderTop().Style != 0 {
		if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style != 0 {
			return PixelType.TopLeftBorder
		} else if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style == 0 {
			return PixelType.TopBorder
		} else if x > w.GetMarginLeft()+1 && x < w.GetMarginLeft()+w.GetWidth() {
			return PixelType.TopBorder
		} else if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style == 0 {
			return PixelType.TopBorder
		} else if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style != 0 {
			return PixelType.TopRightBorder
		}
	} else if y == w.GetMarginTop()+1 && w.GetBorderTop().Style == 0 {
		if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style != 0 {
			return PixelType.LeftBorder
		}
		if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style != 0 {
			return PixelType.RightBorder
		}
	}
	// Bottom border
	if y == w.GetMarginTop()+w.GetHeight() && w.GetBorderBottom().Style != 0 {
		if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style != 0 {
			return PixelType.BottomLeftBorder
		} else if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style == 0 {
			return PixelType.BottomBorder
		} else if x > w.GetMarginLeft()+1 && x < w.GetMarginLeft()+w.GetWidth() {
			return PixelType.BottomBorder
		} else if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style == 0 {
			return PixelType.BottomBorder
		} else if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style != 0 {
			return PixelType.BottomRightBorder
		}
	} else if y == w.GetMarginTop()+w.GetHeight() && w.GetBorderBottom().Style != 1 {
		if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style != 0 {
			return PixelType.LeftBorder
		}
		if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style != 0 {
			return PixelType.RightBorder
		}
	}
	if w.GetMarginTop()+1 < y && y < w.GetMarginTop()+w.GetHeight() {
		if x == w.GetMarginLeft()+1 && w.GetBorderLeft().Style != 0 {
			return PixelType.LeftBorder
		}
		if x == w.GetMarginLeft()+w.GetWidth() && w.GetBorderRight().Style != 0 {
			return PixelType.RightBorder
		}
	}
	// Padding
	var leftBorderWidth, rightBorderWidth, topBorderWidth, bottomBorderWidth int
	if w.GetBorderLeft().Style != 0 {
		leftBorderWidth = 1
	}
	if w.GetBorderRight().Style != 0 {
		rightBorderWidth = 1
	}
	if w.GetBorderTop().Style != 0 {
		topBorderWidth = 1
	}
	if w.GetBorderBottom().Style != 0 {
		bottomBorderWidth = 1
	}
	if y > w.GetMarginTop()+topBorderWidth &&
		y < w.GetMarginTop()+topBorderWidth+w.GetPaddingTop() {
		if w.GetPaddingTop() > 0 {
			return PixelType.Padding
		}
	}
	if x > w.GetMarginTop()+w.GetWidth()-bottomBorderWidth-w.GetPaddingBottom() &&
		x <= w.GetMarginTop()+w.GetWidth()-bottomBorderWidth {
		if w.GetPaddingBottom() > 0 {
			return PixelType.Padding
		}
	}
	if x > w.GetMarginLeft()+leftBorderWidth &&
		x <= w.GetMarginLeft()+leftBorderWidth+w.GetPaddingLeft() {
		if w.GetPaddingLeft() > 0 {
			return PixelType.Padding
		}
	}
	if x > w.GetMarginRight()+w.GetWidth()-rightBorderWidth-w.GetPaddingRight() &&
		x <= w.GetMarginRight()+w.GetWidth()-rightBorderWidth {
		if w.GetPaddingRight() > 0 {
			return PixelType.Padding
		}
	}
	// Content
	return PixelType.Content
}

func (w *Window) Draw() {
	if w.GetScreen() == nil {
		panic("Select's Screen is not defined")
	}
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	for y := 1; y <= w.GetOuterHeight(); y++ {
		for x := 1; x <= w.GetOuterWidth(); x++ {
			w.GetScreen().Screen.SetContent(
				x-1,
				y-1,
				PixelRuneMap[getPositionPixelType(w, x, y)],
				nil,
				borderStyle,
			)
		}
	}
}

func NewWindow() *Window {
	return &Window{
		Container: *NewContainer(),
	}
}

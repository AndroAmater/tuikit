package tuikit

import "github.com/gdamore/tcell/v2"

type Element struct {
	container    ContainsElements
	eventHandler func(event tcell.Event)
	screen       *Screen
	y            int
	x            int
	growX        bool
	growY        bool
	width        int
	height       int
	margin       struct {
		Top    int
		Right  int
		Bottom int
		Left   int
	}
	padding struct {
		Top    int
		Right  int
		Bottom int
		Left   int
	}
	border struct {
		Top    BorderOpts
		Right  BorderOpts
		Bottom BorderOpts
		Left   BorderOpts
	}
}

// IsElement interface

func (e *Element) Draw() {}

// HandlesEvents interface

func (e *Element) HandleEvent(event tcell.Event) {
	e.GetEventHandler()(event)
}

func (e *Element) SetEventHandler(handler func(event tcell.Event)) {
	e.eventHandler = handler
}

func (e *Element) GetEventHandler() func(event tcell.Event) {
	return e.eventHandler
}

// HasContainer interface

func (e *Element) setContainer(c ContainsElements) {
	e.container = c
}

func (e *Element) GetContainer() ContainsElements {
	return e.container
}

// HasPosition interface

func (e *Element) GetX() int {
	return e.x
}

func (e *Element) SetX(x int) {
	e.x = x
}

func (e *Element) GetY() int {
	return e.y
}

func (e *Element) SetY(y int) {
	e.y = y
}

func (e *Element) GetPosition() (int, int) {
	return e.GetX(), e.GetY()
}

func (e *Element) SetPosition(x int, y int) {
	e.SetX(x)
	e.SetY(y)
}

// HasScreen interface

func (e *Element) setScreen(s *Screen) {
	e.screen = s
}

func (e *Element) GetScreen() *Screen {
	return e.screen
}

// HasSize interface

func (e *Element) GetInnerWidth() int {
	var leftBorderWidth, rightBorderWidth int
	if e.GetBorderLeft().Style == BorderStyle.Single {
		leftBorderWidth = 1
	}
	if e.GetBorderRight().Style == BorderStyle.Single {
		rightBorderWidth = 1
	}
	return e.width -
		e.GetPaddingLeft() -
		e.GetPaddingRight() -
		leftBorderWidth -
		rightBorderWidth
}

func (e *Element) GetInnerHeight() int {
	var topBorderHeight, bottomBorderHeight int
	if e.GetBorderTop().Style == BorderStyle.Single {
		topBorderHeight = 1
	}
	if e.GetBorderBottom().Style == BorderStyle.Single {
		bottomBorderHeight = 1
	}
	return e.height -
		e.GetPaddingTop() -
		e.GetPaddingBottom() -
		topBorderHeight -
		bottomBorderHeight
}

func (e *Element) GetInnerSize() (int, int) {
	return e.GetInnerWidth(), e.GetInnerHeight()
}

func (e *Element) GetWidth() int {
	return e.width
}

func (e *Element) GetHeight() int {
	return e.height
}

func (e *Element) GetSize() (int, int) {
	return e.GetWidth(), e.GetHeight()
}

func (e *Element) GetOuterWidth() int {
	return e.GetWidth() + e.GetMarginLeft() + e.GetMarginRight()
}

func (e *Element) GetOuterHeight() int {
	return e.GetHeight() + e.GetMarginTop() + e.GetMarginBottom()
}

func (e *Element) GetOuterSize() (int, int) {
	return e.GetOuterWidth(), e.GetOuterHeight()
}

func (e *Element) SetWidth(width int) {
	e.width = width
}

func (e *Element) SetHeight(height int) {
	e.height = height
}

func (e *Element) SetSize(width int, height int) {
	e.SetWidth(width)
	e.SetHeight(height)
}

func (e *Element) GetGrowX() bool {
	return e.growX
}

func (e *Element) GetGrowY() bool {
	return e.growY
}

func (e *Element) GetGrow() (bool, bool) {
	return e.GetGrowX(), e.GetGrowY()
}

func (e *Element) SetGrowX(growX bool) {
	e.growX = growX
}

func (e *Element) SetGrowY(growY bool) {
	e.growY = growY
}

func (e *Element) SetGrow(growX bool, growY bool) {
	e.SetGrowX(growX)
	e.SetGrowY(growY)
}

func (e *Element) GetMarginTop() int {
	return e.margin.Top
}

func (e *Element) GetMarginRight() int {
	return e.margin.Right
}

func (e *Element) GetMarginBottom() int {
	return e.margin.Bottom
}

func (e *Element) GetMarginLeft() int {
	return e.margin.Left
}

func (e *Element) GetMargin() (int, int, int, int) {
	return e.GetMarginTop(),
		e.GetMarginRight(),
		e.GetMarginBottom(),
		e.GetMarginLeft()
}

func (e *Element) SetMarginTop(margin int) {
	e.margin.Top = margin
}

func (e *Element) SetMarginRight(margin int) {
	e.margin.Right = margin
}

func (e *Element) SetMarginBottom(margin int) {
	e.margin.Bottom = margin
}

func (e *Element) SetMarginLeft(margin int) {
	e.margin.Left = margin
}

func (e *Element) SetMargin(
	marginTop int,
	marginRight int,
	marginBottom int,
	marginLeft int,
) {
	e.SetMarginTop(marginTop)
	e.SetMarginRight(marginRight)
	e.SetMarginBottom(marginBottom)
	e.SetMarginLeft(marginLeft)
}

func (e *Element) GetPaddingTop() int {
	return e.padding.Top
}

func (e *Element) GetPaddingRight() int {
	return e.padding.Right
}

func (e *Element) GetPaddingBottom() int {
	return e.padding.Bottom
}

func (e *Element) GetPaddingLeft() int {
	return e.padding.Left
}

func (e *Element) GetPadding() (int, int, int, int) {
	return e.GetPaddingTop(),
		e.GetPaddingRight(),
		e.GetPaddingBottom(),
		e.GetPaddingLeft()
}

func (e *Element) SetPaddingTop(padding int) {
	e.padding.Top = padding
}

func (e *Element) SetPaddingRight(padding int) {
	e.padding.Right = padding
}

func (e *Element) SetPaddingBottom(padding int) {
	e.padding.Bottom = padding
}

func (e *Element) SetPaddingLeft(padding int) {
	e.padding.Left = padding
}

func (e *Element) SetPadding(
	paddingTop int,
	paddingRight int,
	paddingBottom int,
	paddingLeft int,
) {
	e.SetPaddingTop(paddingTop)
	e.SetPaddingRight(paddingRight)
	e.SetPaddingBottom(paddingBottom)
	e.SetPaddingLeft(paddingLeft)
}

func (e *Element) GetBorderTop() BorderOpts {
	return e.border.Top
}

func (e *Element) GetBorderRight() BorderOpts {
	return e.border.Right
}

func (e *Element) GetBorderBottom() BorderOpts {
	return e.border.Bottom
}

func (e *Element) GetBorderLeft() BorderOpts {
	return e.border.Left
}

func (e *Element) GetBorder() (
	BorderOpts,
	BorderOpts,
	BorderOpts,
	BorderOpts,
) {
	return e.GetBorderTop(),
		e.GetBorderRight(),
		e.GetBorderBottom(),
		e.GetBorderLeft()
}

func (e *Element) SetBorderTop(style BorderStyleType, color string) {
	e.border.Top = BorderOpts{
		Style: style,
		Color: color,
	}
}

func (e *Element) SetBorderRight(style BorderStyleType, color string) {
	e.border.Right = BorderOpts{
		Style: style,
		Color: color,
	}
}

func (e *Element) SetBorderBottom(style BorderStyleType, color string) {
	e.border.Bottom = BorderOpts{
		Style: style,
		Color: color,
	}
}

func (e *Element) SetBorderLeft(style BorderStyleType, color string) {
	e.border.Left = BorderOpts{
		Style: style,
		Color: color,
	}
}

func (e *Element) SetBorder(
	borderTop BorderOpts,
	borderRight BorderOpts,
	borderBottom BorderOpts,
	borderLeft BorderOpts,
) {
	e.SetBorderTop(borderTop.Style, borderTop.Color)
	e.SetBorderRight(borderTop.Style, borderTop.Color)
	e.SetBorderBottom(borderTop.Style, borderTop.Color)
	e.SetBorderLeft(borderTop.Style, borderTop.Color)
}

// Other methods

func NewElement() *Element {
	return &Element{
		container: nil,
	}
}

// Validate interface implementation

var _ IsElement = &Element{}

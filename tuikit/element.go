package tuikit

import (
	"github.com/gdamore/tcell/v2"
	"github.com/google/uuid"
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

type Element struct {
	UUID         uuid.UUID
	container    IsElement
	elements     []IsElement
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
	contentDirection ContainerDirectionType
	contentAlign     ContainerAlignType
	contentJustify   ContainerJustifyType
	content          string
}

func getPositionPixelType(e *Element, x int, y int) PixelTypeType {
	// Margin
	if y <= e.GetMarginTop() {
		return PixelType.Margin
	}
	if y > e.GetMarginTop()+e.GetHeight() {
		return PixelType.Margin
	}
	if x <= e.GetMarginLeft() {
		return PixelType.Margin
	}
	if x > e.GetMarginLeft()+e.GetWidth() {
		return PixelType.Margin
	}
	// Top border
	if y == e.GetMarginTop()+1 && e.GetBorderTop().Style != 0 {
		if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style != 0 {
			return PixelType.TopLeftBorder
		} else if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style == 0 {
			return PixelType.TopBorder
		} else if x > e.GetMarginLeft()+1 && x < e.GetMarginLeft()+e.GetWidth() {
			return PixelType.TopBorder
		} else if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style == 0 {
			return PixelType.TopBorder
		} else if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style != 0 {
			return PixelType.TopRightBorder
		}
	} else if y == e.GetMarginTop()+1 && e.GetBorderTop().Style == 0 {
		if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style != 0 {
			return PixelType.LeftBorder
		}
		if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style != 0 {
			return PixelType.RightBorder
		}
	}
	// Bottom border
	if y == e.GetMarginTop()+e.GetHeight() && e.GetBorderBottom().Style != 0 {
		if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style != 0 {
			return PixelType.BottomLeftBorder
		} else if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style == 0 {
			return PixelType.BottomBorder
		} else if x > e.GetMarginLeft()+1 && x < e.GetMarginLeft()+e.GetWidth() {
			return PixelType.BottomBorder
		} else if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style == 0 {
			return PixelType.BottomBorder
		} else if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style != 0 {
			return PixelType.BottomRightBorder
		}
	} else if y == e.GetMarginTop()+e.GetHeight() && e.GetBorderBottom().Style != 1 {
		if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style != 0 {
			return PixelType.LeftBorder
		}
		if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style != 0 {
			return PixelType.RightBorder
		}
	}
	if e.GetMarginTop()+1 < y && y < e.GetMarginTop()+e.GetHeight() {
		if x == e.GetMarginLeft()+1 && e.GetBorderLeft().Style != 0 {
			return PixelType.LeftBorder
		}
		if x == e.GetMarginLeft()+e.GetWidth() && e.GetBorderRight().Style != 0 {
			return PixelType.RightBorder
		}
	}
	// Padding
	var leftBorderWidth, rightBorderWidth, topBorderWidth, bottomBorderWidth int
	if e.GetBorderLeft().Style != 0 {
		leftBorderWidth = 1
	}
	if e.GetBorderRight().Style != 0 {
		rightBorderWidth = 1
	}
	if e.GetBorderTop().Style != 0 {
		topBorderWidth = 1
	}
	if e.GetBorderBottom().Style != 0 {
		bottomBorderWidth = 1
	}
	if y > e.GetMarginTop()+topBorderWidth &&
		y <= e.GetMarginTop()+topBorderWidth+e.GetPaddingTop() {
		if e.GetPaddingTop() > 0 {
			return PixelType.Padding
		}
	}
	if y > e.GetMarginTop()+e.GetHeight()-bottomBorderWidth-e.GetPaddingBottom() &&
		y <= e.GetMarginTop()+e.GetHeight()-bottomBorderWidth {
		if e.GetPaddingBottom() > 0 {
			return PixelType.Padding
		}
	}
	if x > e.GetMarginLeft()+leftBorderWidth &&
		x <= e.GetMarginLeft()+leftBorderWidth+e.GetPaddingLeft() {
		if e.GetPaddingLeft() > 0 {
			return PixelType.Padding
		}
	}
	if x > e.GetMarginRight()+e.GetWidth()-rightBorderWidth-e.GetPaddingRight() &&
		x <= e.GetMarginRight()+e.GetWidth()-rightBorderWidth {
		if e.GetPaddingRight() > 0 {
			return PixelType.Padding
		}
	}
	// Content
	return PixelType.Content
}

func getPixel(e *Element, p PixelTypeType, x int, y int) rune {
	if p == PixelType.Content {
		topBorderWidth, _, _, rightBorderWidth := e.GetBorderWidth()
		contentXPos := x - e.GetMarginLeft() - e.GetPaddingLeft() - rightBorderWidth
		contentYPos := y - e.GetMarginTop() - e.GetPaddingTop() - topBorderWidth
		contentPos := contentXPos + ((contentYPos - 1) * e.GetInnerWidth())
		if len(e.GetContent()) < contentPos {
			return PixelRuneMap[p]
		}
		return rune(e.GetContent()[contentPos-1])
	}
	return PixelRuneMap[p]
}

// IsElement interface

func (e *Element) GetUUID() uuid.UUID {
	if e.UUID == uuid.Nil {
		e.UUID = uuid.New()
	}
	return e.UUID
}

func (e *Element) Draw() {
	if e.GetScreen() == nil {
		panic("Element's Screen is not defined")
	}
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	type pixel struct {
		x         int
		y         int
		pixelType PixelTypeType
	}
	ch := make(chan pixel)
	getRune := func(e *Element, x int, y int) {
		ch <- pixel{
			x:         x,
			y:         y,
			pixelType: getPositionPixelType(e, x, y),
		}
	}
	for y := 1; y <= e.GetOuterHeight(); y++ {
		for x := 1; x <= e.GetOuterWidth(); x++ {
			go getRune(e, x, y)
		}
	}
	for i := 0; i < e.GetOuterHeight()*e.GetOuterWidth(); i++ {
		p := <-ch
		e.GetScreen().Screen.SetContent(
			p.x+e.GetAbsoluteX(),
			p.y+e.GetAbsoluteY(),
			getPixel(e, p.pixelType, p.x, p.y),
			nil,
			borderStyle,
		)
	}
	e.DrawChildren()
}

// HandlesEvents interface

func (e *Element) HandleEvent(event tcell.Event) {
	for _, e := range e.elements {
		e.HandleEvent(event)
	}
	e.eventHandler(event)
}

func (e *Element) SetEventHandler(handler func(event tcell.Event)) {
	e.eventHandler = handler
}

func (e *Element) GetEventHandler() func(event tcell.Event) {
	return e.eventHandler
}

// HasContainer interface

func (e *Element) setContainer(c IsElement) {
	e.container = c
}

func (e *Element) GetContainer() IsElement {
	return e.container
}

// HasPosition interface

func (e *Element) GetX() int {
	return e.x
}

func (e *Element) GetY() int {
	return e.y
}

// TODO: Take align and justify into account
func (e *Element) GetAbsoluteX() int {
	if e.GetContainer() == nil {
		return e.GetX()
	}
	_, _, _, leftBorderWidth := e.GetContainer().GetBorderWidth()
	offset := 0
	if e.GetContainer().GetContentDirection() == ContainerDirection.Row {
		for _, child := range e.GetContainer().GetChildren() {
			if child.GetUUID() == e.GetUUID() {
				break
			}
			offset += child.GetOuterWidth()
		}
	}
	return e.GetX() +
		e.GetContainer().GetAbsoluteX() +
		e.GetContainer().GetMarginLeft() +
		e.GetContainer().GetPaddingLeft() +
		leftBorderWidth +
		offset
}

// TODO: Take align and justify into account
func (e *Element) GetAbsoluteY() int {
	if e.GetContainer() == nil {
		return e.GetY()
	}
	topBorderWidth, _, _, _ := e.GetContainer().GetBorderWidth()
	offset := 0
	if e.GetContainer().GetContentDirection() == ContainerDirection.Column {
		for _, child := range e.GetContainer().GetChildren() {
			if child.GetUUID() == e.GetUUID() {
				break
			}
			offset += child.GetOuterHeight()
		}
	}
	return e.GetY() +
		e.GetContainer().GetAbsoluteY() +
		e.GetContainer().GetMarginTop() +
		e.GetContainer().GetPaddingTop() +
		topBorderWidth +
		offset
}

func (e *Element) GetPosition() (int, int) {
	return e.GetX(), e.GetY()
}

func (e *Element) GetAbsolutePosition() (int, int) {
	return e.GetAbsoluteX(), e.GetAbsoluteY()
}

func (e *Element) SetX(x int) {
	e.x = x
}

func (e *Element) SetY(y int) {
	e.y = y
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
	if e.screen != nil {
		return e.screen
	}
	return e.GetContainer().GetScreen()
}

// HasSize interface

func (e *Element) GetInnerWidth() int {
	_, rightBorderWidth, _, leftBorderWidth := e.GetBorderWidth()
	return e.GetWidth() -
		e.GetPaddingLeft() -
		e.GetPaddingRight() -
		leftBorderWidth -
		rightBorderWidth
}

func (e *Element) GetInnerHeight() int {
	topBorderHeight, _, bottomBorderHeight, _ := e.GetBorderWidth()
	return e.GetHeight() -
		e.GetPaddingTop() -
		e.GetPaddingBottom() -
		topBorderHeight -
		bottomBorderHeight
}

func (e *Element) GetInnerSize() (int, int) {
	return e.GetInnerWidth(), e.GetInnerHeight()
}

func (e *Element) GetWidth() int {
	if e.growX {
		if e.GetContainer() != nil {
			width := e.GetContainer().GetInnerWidth() -
				e.GetMarginLeft() -
				e.GetMarginRight()
			if e.GetContainer().GetContentDirection() == ContainerDirection.Row {
				reservedWidth := 0
				for _, child := range e.GetContainer().GetXNonGrowableChildren() {
					if child == e {
						continue
					}
					reservedWidth += child.GetOuterWidth()
				}
				for _, child := range e.GetContainer().GetXGrowableChildren() {
					if child == e {
						continue
					}
					reservedWidth += child.GetMarginRight()
					reservedWidth += child.GetMarginLeft()
				}
				return (width - reservedWidth) /
					len(e.GetContainer().GetXGrowableChildren())
			}
			return width
		}
		if e.GetScreen() == nil {
			panic("Element's screen is not defined")
		}
		return e.GetScreen().GetWidth() -
			e.GetMarginLeft() -
			e.GetMarginRight()
	}
	return e.width
}

func (e *Element) GetHeight() int {
	if e.growY {
		if e.GetContainer() != nil {
			height := e.GetContainer().GetInnerHeight() -
				e.GetMarginTop() -
				e.GetMarginBottom()
			if e.GetContainer().GetContentDirection() == ContainerDirection.Column {
				reservedHeight := 0
				for _, child := range e.GetContainer().GetYNonGrowableChildren() {
					if child.GetUUID() == e.GetUUID() {
						continue
					}
					reservedHeight += child.GetOuterHeight()
				}
				for _, child := range e.GetContainer().GetYGrowableChildren() {
					if child.GetUUID() == e.GetUUID() {
						continue
					}
					reservedHeight += child.GetMarginTop()
					reservedHeight += child.GetMarginBottom()
				}
				return (height - reservedHeight) /
					len(e.GetContainer().GetYGrowableChildren())
			}
			return height
		}
		if e.GetScreen() == nil {
			panic("Element's screen is not defined")
		}
		return e.GetScreen().GetHeight() -
			e.GetMarginTop() -
			e.GetMarginBottom()
	}
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

func (e *Element) GetBorderWidth() (int, int, int, int) {
	var leftBorderWidth, rightBorderWidth, topBorderHeight, bottomBorderHeight int
	if e.GetBorderLeft().Style == BorderStyle.Single {
		leftBorderWidth = 1
	}
	if e.GetBorderRight().Style == BorderStyle.Single {
		rightBorderWidth = 1
	}
	if e.GetBorderTop().Style == BorderStyle.Single {
		topBorderHeight = 1
	}
	if e.GetBorderBottom().Style == BorderStyle.Single {
		bottomBorderHeight = 1
	}
	return topBorderHeight,
		rightBorderWidth,
		bottomBorderHeight,
		leftBorderWidth
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

func (e *Element) SetFullBorder(style BorderStyleType, color string) {
	e.SetBorderTop(style, color)
	e.SetBorderRight(style, color)
	e.SetBorderBottom(style, color)
	e.SetBorderLeft(style, color)
}

// ContainsElements interface

func (e *Element) AddChild(el IsElement) {
	e.elements = append(e.elements, el)
	el.setContainer(e)
}

func (e *Element) GetChildren() []IsElement {
	return e.elements
}

func (e *Element) DrawChildren() {
	for _, el := range e.elements {
		el.Draw()
	}
}

func (e *Element) GetContentDirection() ContainerDirectionType {
	return e.contentDirection
}

func (e *Element) GetContentAlign() ContainerAlignType {
	return e.contentAlign
}

func (e *Element) GetContentJustify() ContainerJustifyType {
	return e.contentJustify
}

func (e *Element) SetContentDirection(direction ContainerDirectionType) {
	e.contentDirection = direction
}

func (e *Element) SetContentAlign(align ContainerAlignType) {
	e.contentAlign = align
}

func (e *Element) SetContentJustify(justify ContainerJustifyType) {
	e.contentJustify = justify
}

func (e *Element) GetXGrowableChildren() []IsElement {
	var growable []IsElement
	for _, el := range e.GetChildren() {
		if el.GetGrowX() == true {
			growable = append(growable, el)
		}
	}
	return growable
}

func (e *Element) GetYGrowableChildren() []IsElement {
	var growable []IsElement
	for _, el := range e.GetChildren() {
		if el.GetGrowY() == true {
			growable = append(growable, el)
		}
	}
	return growable
}

func (e *Element) GetXNonGrowableChildren() []IsElement {
	var growable []IsElement
	for _, el := range e.GetChildren() {
		if el.GetGrowX() != true {
			growable = append(growable, el)
		}
	}
	return growable
}

func (e *Element) GetYNonGrowableChildren() []IsElement {
	var growable []IsElement
	for _, el := range e.GetChildren() {
		if el.GetGrowY() != true {
			growable = append(growable, el)
		}
	}
	return growable
}

func (e *Element) RemoveAllChildren() {
	e.elements = []IsElement{}
}

// HasContent interface

func (e *Element) GetContent() string {
	return e.content
}

func (e *Element) SetContent(content string) {
	e.content = content
}

// Other methods

func NewElement() *Element {
	return &Element{
		UUID:         uuid.New(),
		eventHandler: func(event tcell.Event) {},
	}
}

// Validate interface implementation

var _ IsElement = &Element{}

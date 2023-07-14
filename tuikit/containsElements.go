package tuikit

type (
	ContainerDirectionType int
	ContainerAlignType     int
	ContainerJustifyType   int
)

const (
	dirRow ContainerDirectionType = iota
	dirColumn
)

// Border styles enum
var ContainerDirection = struct {
	Row    ContainerDirectionType
	Column ContainerDirectionType
}{
	Row:    dirRow,
	Column: dirColumn,
}

const (
	alignStart ContainerAlignType = iota
	alignCenter
	alignEnd
)

var ContainerAlign = struct {
	Start  ContainerAlignType
	Center ContainerAlignType
	End    ContainerAlignType
}{
	Start:  alignStart,
	Center: alignCenter,
	End:    alignEnd,
}

const (
	justifyStart ContainerJustifyType = iota
	justifyCenter
	justifyEnd
	justifySpaceBetween
	justifySpaceAround
)

var ContainerJustify = struct {
	Start        ContainerJustifyType
	Center       ContainerJustifyType
	End          ContainerJustifyType
	SpaceBetween ContainerJustifyType
	SpaceAround  ContainerJustifyType
}{
	Start:        justifyStart,
	Center:       justifyCenter,
	End:          justifyEnd,
	SpaceBetween: justifySpaceBetween,
	SpaceAround:  justifySpaceAround,
}

type ContainsElements interface {
	// Adds an element to the container
	AddChild(e IsElement)
	// Removes an element from the container
	GetChildren() []IsElement
	// Draws the container's children
	DrawChildren()
	// Gets the container's content direction
	GetContentDirection() ContainerDirectionType
	// Gets the container's content justify
	GetContentJustify() ContainerJustifyType
	// Gets the container's content align
	GetContentAlign() ContainerAlignType
	// Sets the container's content direction
	SetContentDirection(ContainerDirectionType)
	// Sets the container's content justify
	SetContentJustify(ContainerJustifyType)
	// Sets the container's content align
	SetContentAlign(ContainerAlignType)
	// Gets ther container's X growable children
	GetXGrowableChildren() []IsElement
	// Gets ther container's Y growable children
	GetYGrowableChildren() []IsElement
	// Gets ther container's X non-growable children
	GetXNonGrowableChildren() []IsElement
	// Gets ther container's Y non-growable children
	GetYNonGrowableChildren() []IsElement
	// Removes all element's children
	RemoveAllChildren()
}

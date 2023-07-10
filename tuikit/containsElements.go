package tuikit

type ContainsElements interface {
	AddElement(e IsElement)
	GetElements() []IsElement
	DrawChildren()
}

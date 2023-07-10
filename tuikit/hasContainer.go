package tuikit

type HasContainer interface {
	setContainer(c ContainsElements)
	GetContainer() ContainsElements
}

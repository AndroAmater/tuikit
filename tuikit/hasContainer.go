package tuikit

type HasContainer interface {
	// Sets the element's container
	setContainer(c ContainsElements)
	// Gets the element's container
	GetContainer() ContainsElements
}

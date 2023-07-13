package tuikit

type HasContainer interface {
	// Sets the element's container
	setContainer(c IsElement)
	// Gets the element's container
	GetContainer() IsElement
}

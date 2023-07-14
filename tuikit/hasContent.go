package tuikit

type HasContent interface {
	// Gets the element's content
	GetContent() string
	// Sets the element's content
	SetContent(content string)
}

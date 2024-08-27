package browser

import "fmt"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) BrowserClose() error {
	// Implement browser closing logic here

	fmt.Println("Browser Close")

	return fmt.Errorf("TEST Error Browser Close")
}

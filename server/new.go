package server

import "fmt"

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) ServerClose() error {
	// Implement server closing logic here
	fmt.Println("Server Close")

	return fmt.Errorf("TEST Error Server Close")
}

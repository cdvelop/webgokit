package compiler

type consoleAdapter interface {
	Info(message string)
	Error(error)
}

type handler struct {
	consoleAdapter
}

func New(ca consoleAdapter) *handler {
	return &handler{
		consoleAdapter: ca,
	}
}

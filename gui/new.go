package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type compilerAdapter interface {
	InputDirectoryChange(newDirectory string)
	OutputDirectoryChange(newDirectory string)
}

type browserAdapter interface {
	BrowserClose() error
}

type serverAdapter interface {
	ServerClose() error
}

type consoleAdapter interface {
	Console() fyne.CanvasObject
	Info(message string)
	Success(message string)
	Error(message string)
}

type handler struct {
	// console *widget.Entry
	window fyne.Window
	consoleAdapter

	compiler compilerAdapter
	browser  browserAdapter
	server   serverAdapter
}

func New(cons consoleAdapter, c compilerAdapter, b browserAdapter, s serverAdapter) *handler {

	a := app.NewWithID("com.webgokit.cdvelop.github")
	a.Settings().SetTheme(newFysionTheme())

	h := &handler{
		window:         a.NewWindow("WebGoKit"),
		consoleAdapter: cons,

		compiler: c,
		browser:  b,
		server:   s,
	}

	return h
}

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type browserAdapter interface {
	BrowserClose() error
}

type serverAdapter interface {
	ServerClose() error
}

type handler struct {
	// console *widget.Entry
	window fyne.Window
	*Console
	logFilePath string // ej: log.txt

	browser browserAdapter

	server serverAdapter
}

func New(b browserAdapter, s serverAdapter) *handler {

	a := app.NewWithID("com.webgokit.cdvelop.github")
	a.Settings().SetTheme(newFysionTheme())

	h := &handler{
		window:      a.NewWindow("WebGoKit"),
		Console:     NewConsole(),
		logFilePath: "./log.txt",
		browser:     b,
		server:      s,
	}

	return h
}

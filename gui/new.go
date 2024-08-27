package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type handler struct {
	// console *widget.Entry
	window fyne.Window
	*Console
}

func New() *handler {

	a := app.NewWithID("com.webgokit.cdvelop.github")
	a.Settings().SetTheme(newFysionTheme())

	h := &handler{
		window:  a.NewWindow("WebGoKit"),
		Console: NewConsole(),
	}

	return h
}

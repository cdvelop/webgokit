package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Layout izquierdo
func (h *handler) formLeftContent() *fyne.Container {

	// URL LOGIN
	urlLoginLabel := widget.NewLabel("URL Login")
	entryLoginURL := widget.NewEntry()
	// entryLoginURL.SetText(h.url.login)
	// entryLoginURL.SetPlaceHolder(h.url.login)

	// URL HOME
	urlHomeLabel := widget.NewLabel("URL Home")
	entryHomeURL := widget.NewEntry()
	// entryHomeURL.SetText(h.url.home)
	// entryHomeURL.SetPlaceHolder(h.url.home)
	entryHomeURL.OnChanged = func(text string) {

		// h.url.login = entryHomeURL.Text + "/login"
		// entryLoginURL.SetPlaceHolder(h.url.login)
		// entryLoginURL.SetText(h.url.login)

	}

	// Botón de procesamiento
	processButton := widget.NewButton("Procesar", func() {

		h.logToConsole("Procesando...\n\n")

		h.logToConsole("\n\n")
	})

	return container.NewVBox(
		urlHomeLabel,
		entryHomeURL,
		urlLoginLabel,
		entryLoginURL,
		layout.NewSpacer(),
		processButton,
	)
}

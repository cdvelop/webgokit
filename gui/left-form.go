package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) formLeftContent() *fyne.Container {

	// Botón de procesamiento
	processButton := widget.NewButton("Procesar", func() {
		h.Info("Procesando...")
		// Aquí puedes agregar la lógica de procesamiento usando
		// projectPathEntry.Text y compilationPathEntry.Text
	})

	// Crear el acordeón con los dos items
	accordion := widget.NewAccordion(h.makeDevelopForm(), h.makeDeployForm())

	// Expandir el primer item por defecto
	accordion.Open(0)

	return container.NewVBox(
		accordion,
		processButton,
	)
}

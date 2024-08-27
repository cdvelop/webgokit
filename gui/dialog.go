package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) customConfirmDialog(message string, callback func(bool)) {
	var dialog *widget.PopUp

	content := container.NewVBox(
		widget.NewLabel(message),
		container.NewHBox(
			widget.NewButton("No", func() {
				dialog.Hide()
				callback(false)
			}),
			widget.NewButton("Sí", func() {
				dialog.Hide()
				callback(true)
			}),
		),
	)

	dialog = widget.NewModalPopUp(content, h.window.Canvas())
	dialog.Show()
}

func (h *handler) showFolderDialog(entry *widget.Entry, successMsg string, callback func()) {
	// Crear el cuadro de diálogo para seleccionar la carpeta
	folderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			h.Error(err.Error())
			return
		}
		if uri == nil {
			return
		}

		entry.SetText(uri.Path())
		h.Success(successMsg)

		callback()
	}, h.window)

	// Establecer el tamaño del cuadro de diálogo
	folderDialog.Resize(fyne.NewSize(800, 600)) // Cambia el tamaño según tus necesidades

	// Mostrar el diálogo personalizado
	folderDialog.Show()
}

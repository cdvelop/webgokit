package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) makeDevelopForm() *widget.AccordionItem {
	// FOLDER PROJECT PATH
	inPathEntry := widget.NewEntry()
	inPathButton := widget.NewButton("Select In Folder", func() {
		// Crear el cuadro de diálogo para seleccionar la carpeta
		folderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				h.Error(err.Error())
				return
			}
			if uri == nil {
				return
			}
			inPathEntry.SetText(uri.Path())
			h.Success("Module folder selected")
		}, h.window)

		// Establecer el tamaño del cuadro de diálogo
		folderDialog.Resize(fyne.NewSize(800, 600)) // Cambia el tamaño según tus necesidades

		// Mostrar el diálogo personalizado
		folderDialog.Show()
	})

	// DESTINATION COMPILER PROJECT
	outPathEntry := widget.NewEntry()
	outPathButton := widget.NewButton("Select Out Folder", func() {
		// Crear el cuadro de diálogo para seleccionar la carpeta
		folderDialog := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				h.Error(err.Error())
				return
			}
			if uri == nil {
				return
			}
			outPathEntry.SetText(uri.Path())
		}, h.window)

		// Establecer el tamaño del cuadro de diálogo
		folderDialog.Resize(fyne.NewSize(800, 600)) // Cambia el tamaño según tus necesidades

		// Mostrar el diálogo personalizado
		folderDialog.Show()
	})

	moduleDevelop := widget.NewAccordionItem("Develop", container.NewVBox(
		inPathEntry,
		inPathButton,
		outPathEntry,
		outPathButton,
	))

	return moduleDevelop
}

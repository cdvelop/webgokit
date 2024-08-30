package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) makeDevelopForm() *widget.AccordionItem {
	// FOLDER PROJECT PATH
	inPathEntry := widget.NewEntry()
	folderIN := h.conf.Get().FolderInput()
	if folderIN != "" {
		inPathEntry.SetText(folderIN)
		h.watcher.RegisterFiles(folderIN)
	}

	inPathButton := widget.NewButton("Select In Folder", func() {
		h.showFolderDialog(inPathEntry, "Folder IN", func() {
			folderIN = inPathEntry.Text
			// notificar al compilador el cambio
			h.compiler.InputDirectoryChange(folderIN)
			// registrar archivos nuevo directorio
			h.watcher.RegisterFiles(folderIN)
			// guardar nueva configuración
			h.conf.Set().FolderInput(folderIN)

		})

	})

	// DESTINATION COMPILER PROJECT
	outPathEntry := widget.NewEntry()
	folderOUT := h.conf.Get().FolderOutput()
	if folderOUT != "" {
		outPathEntry.SetText(folderOUT)
	}
	outPathButton := widget.NewButton("Select Out Folder", func() {
		h.showFolderDialog(outPathEntry, "Folder OUT", func() {
			folderOUT = outPathEntry.Text
			// notificar al compilador
			h.compiler.OutputDirectoryChange(folderOUT)

			// guardar nueva configuración
			h.conf.Set().FolderOutput(folderOUT)
		})
	})

	moduleDevelop := widget.NewAccordionItem("Develop", container.NewVBox(
		inPathEntry,
		inPathButton,
		outPathEntry,
		outPathButton,
	))

	return moduleDevelop
}

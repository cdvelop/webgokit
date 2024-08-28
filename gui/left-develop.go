package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) makeDevelopForm() *widget.AccordionItem {
	// FOLDER PROJECT PATH
	inPathEntry := widget.NewEntry()
	inPathButton := widget.NewButton("Select In Folder", func() {
		h.showFolderDialog(inPathEntry, "Folder IN", func() {

			h.compiler.InputDirectoryChange(inPathEntry.Text)
			// registrar archivos nuevo directorio
			h.watcher.RegisterFiles(inPathEntry.Text)
		})

	})

	// DESTINATION COMPILER PROJECT
	outPathEntry := widget.NewEntry()
	outPathButton := widget.NewButton("Select Out Folder", func() {
		h.showFolderDialog(outPathEntry, "Folder OUT", func() {

			h.compiler.OutputDirectoryChange(outPathEntry.Text)
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

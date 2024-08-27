package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) showConfirmDialog(title, message string, callback func(bool)) {
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

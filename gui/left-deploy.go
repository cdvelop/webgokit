package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) makeDeployForm() *widget.AccordionItem {

	deployProject := widget.NewAccordionItem("Deploy", container.NewVBox(
		widget.NewLabel("Deploy PROJECT"),
	))

	return deployProject
}

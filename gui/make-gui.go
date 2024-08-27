package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (h *handler) makeGUI() fyne.CanvasObject {

	top := h.makeBanner()
	left := container.NewPadded(container.NewVBox(h.formLeftContent()))
	content := container.NewPadded(container.NewScroll(h.Console()))

	// return container.NewBorder(h.makeBanner(),nil,container.NewVBox(h.formLeftContent()),nil,container.NewScroll(h.Console.scroll),)

	dividers := [2]fyne.CanvasObject{
		widget.NewSeparator(), widget.NewSeparator(),
	}

	objs := []fyne.CanvasObject{content, top, left, dividers[0], dividers[1]}

	return container.New(h.newFusionLayout(top, left, nil, content, dividers), objs...)

}

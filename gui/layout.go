package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWidth = 350 // ancho la barra lateral

type fusionLayout struct {
	top, left, right, content fyne.CanvasObject
	dividers                  [2]fyne.CanvasObject
}

func (h *handler) newFusionLayout(top, left, right, content fyne.CanvasObject, dividers [2]fyne.CanvasObject) fyne.Layout {

	return &fusionLayout{top, left, right, content, dividers}

}

func (l *fusionLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	// ubicar barra lateral
	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(fyne.NewSize(sideWidth, size.Height-topHeight))

	// ubicar contenido
	l.content.Move(fyne.NewPos(sideWidth, topHeight))
	l.content.Resize(fyne.NewSize(size.Width-sideWidth, size.Height-topHeight))

	// log.Println("Size:", size)
	// ubicar lineas divisoras
	// dividerThickness := float32(5)
	dividerThickness := theme.SeparatorThicknessSize()
	// ubicar linea divisora
	l.dividers[0].Move(fyne.NewPos(0, topHeight))
	l.dividers[0].Resize(fyne.NewSize(size.Width, dividerThickness))

	l.dividers[1].Move(fyne.NewPos(sideWidth, topHeight))
	l.dividers[1].Resize(fyne.NewSize(dividerThickness, size.Height-topHeight))
}

func (l *fusionLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {

	borders := fyne.NewSize(
		sideWidth*3,
		l.top.MinSize().Height,
	)

	return borders.AddWidthHeight(100, 100)

	// ajustar el tamaño mínimo al iniciar
	// return fyne.Size{Width: 1080, Height: 720}
}

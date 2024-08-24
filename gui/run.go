package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	h := &handler{
		console: widget.NewMultiLineEntry(),
	}

	a := app.New()
	w := a.NewWindow("WebGoKit")

	// Asigna el icono a la aplicación
	// icon := fyne.NewStaticResource("icon", iconData)
	// myApp.SetIcon(icon)

	// configuración Consola
	h.console.SetMinRowsVisible(10)
	h.console.Wrapping = fyne.TextWrapWord
	// Redirigir fmt.Print y log a la consola
	// log.SetOutput(h)

	// Contenedor principal
	content := container.NewHSplit(
		container.NewVBox(h.formLeftContent()),
		container.NewScroll(h.console),
	)

	content.Offset = 0.3

	w.SetContent(content)

	// w.SetFixedSize(true)

	// w.Resize(fyne.NewSize(100, 100))
	w.Resize(fyne.NewSize(1080, 720))

	w.ShowAndRun()
}

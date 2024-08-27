//go:generate fyne bundle -o bundled.go assets
package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func Run() {

	a := app.NewWithID("com.webgokit.cdvelop.github")
	a.Settings().SetTheme(newFysionTheme())
	w := a.NewWindow("WebGoKit")

	h := &handler{
		window:  w,
		Console: NewConsole(),
	}

	w.Resize(fyne.NewSize(1080, 720))

	w.SetContent(h.makeGUI())
	w.ShowAndRun()
}

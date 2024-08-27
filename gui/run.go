//go:generate fyne bundle -o bundled.go assets
package gui

import (
	"time"

	"fyne.io/fyne/v2"
)

func (h *handler) Run() {

	h.window.Resize(fyne.NewSize(1080, 720))

	h.window.SetContent(h.makeGUI())

	// Configurar el manejador de cierre
	h.window.SetCloseIntercept(func() {
		// Aquí puedes realizar cualquier limpieza o guardar el estado antes de cerrar
		h.showConfirmDialog("¿ Deseas cerrar la aplicación ?", func(ok bool) {
			if ok {
				// Ejecuta aquí cualquier función que necesites antes de cerrar

				if err := h.browser.BrowserClose(); err != nil {
					h.Error(err.Error())
				}
				if err := h.server.ServerClose(); err != nil {
					h.Error(err.Error())
				}

				// esperar x segundo para que se cierre
				time.Sleep(1 * time.Second)

				// cerrar la aplicación
				h.window.Close()
			}
		})
	})

	h.window.ShowAndRun()
}

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
		h.customConfirmDialog("¿ Deseas cerrar la aplicación ?", func(ok bool) {
			if ok {
				// Ejecuta aquí cualquier función que necesites antes de cerrar

				if err := h.browser.BrowserClose(); err != nil {
					h.Error(err)
				}
				if err := h.server.ServerClose(); err != nil {
					h.Error(err)
				}

				// esperar x tiempo para el cierre
				time.Sleep(200 * time.Millisecond)

				// cerrar la aplicación
				h.window.Close()
			}
		})
	})

	go h.watcher.FileWatcherStart()

	h.window.ShowAndRun()
}

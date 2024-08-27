//go:generate fyne bundle -o bundled.go assets
package gui

import (
	"fyne.io/fyne/v2"
)

func (h *handler) Run() {

	h.window.Resize(fyne.NewSize(1080, 720))

	h.window.SetContent(h.makeGUI())

	// Configurar el manejador de cierre
	h.window.SetCloseIntercept(func() {
		// Aquí puedes realizar cualquier limpieza o guardar el estado antes de cerrar
		h.showConfirmDialog("Confirmar cierre", "¿ Deseas cerrar la aplicación ?", func(ok bool) {
			if ok {
				// Ejecuta aquí cualquier función que necesites antes de cerrar
				// Por ejemplo, podrías llamar a una función de limpieza:
				// handler.CleanupBeforeExit()
				h.window.Close()
			}
		})
	})

	h.window.ShowAndRun()
}

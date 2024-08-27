package console

import "fyne.io/fyne/v2"

func (c *handler) Console() fyne.CanvasObject {
	return c.scroll
}

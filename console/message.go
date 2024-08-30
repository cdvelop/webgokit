package console

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2/canvas"
)

func (c *handler) Info(message string) {
	c.addText(message, c.colorNormal)
}
func (c *handler) Success(message string) {
	c.addText(message, c.colorGreen)
}
func (c *handler) Debug(message string) {
	c.addText(message, c.colorBlue)
}

func (c *handler) Error(e error) {
	if e != nil {
		// mostrar mensaje en consola gui
		c.addText(e.Error(), c.colorRed)
		// guardar mensaje en archivo log.txt
		log.Println(e.Error())
	}
}

func (c *handler) addText(text string, col color.Color) {
	if text != "" {
		t := canvas.NewText(text, col)
		c.texts = append(c.texts, t)
		c.content.Add(t)
		c.scroll.ScrollToBottom()
	}
}

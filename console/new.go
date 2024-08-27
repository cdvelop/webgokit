package console

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type handler struct {
	logFilePath string // ej: log.txt
	texts       []*canvas.Text
	content     *fyne.Container
	scroll      *container.Scroll

	colorNormal color.Color
	colorGreen  color.Color
	colorBlue   color.Color
	colorRed    color.Color
}

func New() *handler {
	c := &handler{
		logFilePath: "./log.txt",
		texts:       make([]*canvas.Text, 0),
		content:     container.New(layout.NewVBoxLayout()),

		// colorNormal: theme.PrimaryColor(),
		colorNormal: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		colorGreen:  color.RGBA{R: 0, G: 100, B: 0, A: 255},
		colorBlue:   color.RGBA{R: 0, G: 0, B: 255, A: 255},
		colorRed:    color.RGBA{R: 255, A: 255},
	}
	c.scroll = container.NewScroll(c.content)

	log.SetOutput(c)

	return c
}

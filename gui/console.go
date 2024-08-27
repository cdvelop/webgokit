package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Console struct {
	texts   []*canvas.Text
	content *fyne.Container
	scroll  *container.Scroll

	colorNormal color.Color
	colorGreen  color.Color
	colorBlue   color.Color
	colorRed    color.Color
}

func NewConsole() *Console {
	c := &Console{
		texts:   make([]*canvas.Text, 0),
		content: container.New(layout.NewVBoxLayout()),

		// colorNormal: theme.PrimaryColor(),
		colorNormal: color.RGBA{R: 0, G: 0, B: 0, A: 255},
		colorGreen:  color.RGBA{R: 0, G: 100, B: 0, A: 255},
		colorBlue:   color.RGBA{R: 0, G: 0, B: 255, A: 255},
		colorRed:    color.RGBA{R: 255, A: 255},
	}
	c.scroll = container.NewScroll(c.content)

	// log.SetOutput(c)

	return c
}

func (c *Console) Info(message string) {
	c.addText(message, c.colorNormal)
}
func (c *Console) Success(message string) {
	c.addText(message, c.colorGreen)
}
func (c *Console) Debug(message string) {
	c.addText(message, c.colorBlue)
}

func (c *Console) Error(message string) {
	c.addText(message, c.colorRed)
}

func (c *Console) addText(text string, col color.Color) {
	t := canvas.NewText(text, col)
	c.texts = append(c.texts, t)
	c.content.Add(t)
	c.scroll.ScrollToBottom()
}

func (c *Console) Write(p []byte) (n int, err error) {
	c.Error(string(p))

	return len(p), nil
}

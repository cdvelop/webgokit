package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fysionTheme struct {
	fyne.Theme
}

func newFysionTheme() fyne.Theme {

	return &fysionTheme{
		Theme: theme.DefaultTheme()}

}

func (t *fysionTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	// log.Println("COLOR NAME:", name)
	// return t.Theme.Color(name, theme.VariantDark)
	return t.Theme.Color(name, variant)
}

func (t *fysionTheme) Size(name fyne.ThemeSizeName) float32 {

	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}

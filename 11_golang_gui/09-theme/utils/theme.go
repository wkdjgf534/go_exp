package utils

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct {
	FontSize float32
}

// Color
func (c *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{30, 30, 30, 255}
	case theme.ColorNameButton:
		return color.RGBA{0, 122, 255, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{255, 215, 0, 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

// Font
func (c *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

// Icon
func (c *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// Size
func (c *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return c.FontSize
	default:
		return theme.DefaultTheme().Size(name)
	}
}

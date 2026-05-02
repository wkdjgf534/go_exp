package main

import (
	"09-theme/utils"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ThemeVariant struct {
	fyne.Theme
	variant fyne.ThemeVariant
}

func (f *ThemeVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.variant)
}

func main() {
	app := app.NewWithID("themes.com")
	window := app.NewWindow("themes Showcase")
	window.Resize(fyne.NewSize(500, 500))

	// load user preferences
	currentFontSize := app.Preferences().FloatWithFallback("font_size", 14)
	currentTheme := app.Preferences().StringWithFallback("theme", "custom")
	customTheme := &utils.CustomTheme{FontSize: float32(currentFontSize)}

	switch currentTheme {
	case "light":
		app.Settings().SetTheme(&ThemeVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
	case "dark":
		app.Settings().SetTheme(&ThemeVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
	default:
		app.Settings().SetTheme(customTheme)
	}

	// display the settings
	showSettingsButton := func() {
		themeSelect := widget.NewSelect([]string{"light", "dark", "custom"}, func(selected string) {
			switch selected {
			case "light":
				app.Settings().SetTheme(&ThemeVariant{Theme: theme.DefaultTheme(), variant: theme.VariantLight})
				app.Preferences().SetString("theme", "light")
			case "dark":
				app.Settings().SetTheme(&ThemeVariant{Theme: theme.DefaultTheme(), variant: theme.VariantDark})
				app.Preferences().SetString("theme", "dark")
			default:
				app.Settings().SetTheme(customTheme)
				app.Preferences().SetString("theme", "custom")
			}
		})
		themeSelect.SetSelected(currentTheme)

		// font size slider
		fontSizeSlider := widget.NewSlider(10, 24)
		fontSizeSlider.SetValue(currentFontSize)

		fontSizeLabel := widget.NewLabel(fmt.Sprintf("Font Size: %.0f", currentFontSize))
		fontSizeSlider.OnChanged = func(size float64) {
			customTheme.FontSize = float32(size)
			app.Settings().SetTheme(customTheme)
			fontSizeLabel.SetText(fmt.Sprintf("Font Size: %.0f", size))
			app.Preferences().SetFloat("font_size", float64(size))
		}

		content := container.NewGridWithColumns(
			2,
			widget.NewLabel("Theme:"),
			themeSelect,
			widget.NewLabel("Font Size:"),
			fontSizeSlider,
			fontSizeLabel,
		)

		dialog.ShowCustom("Settings", "Close", content, window)
	}

	settingButton := widget.NewButton("Settings", func() {
		showSettingsButton()
	})

	window.SetContent(container.NewVBox(settingButton))
	window.ShowAndRun()
}

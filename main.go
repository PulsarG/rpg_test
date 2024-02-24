package main

import (
	//"context"
	//"log"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/dialog"
	//"fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	App := app.New()
	mainWindow := App.NewWindow("RPG TEST")
	mainWindow.Resize(fyne.NewSize(500, 500))
	/* positiv := 1
	negativ := -1 */

	mainCount := 0
	infoPanel := widget.NewLabel(strconv.Itoa(mainCount))

	stopButton := widget.NewButton("Минус", func() { plusOrMinus(&mainCount, -1, infoPanel) })

	startButton := widget.NewButton("Плюс", func() { plusOrMinus(&mainCount, 1, infoPanel) })

	cont := container.NewVBox(infoPanel, startButton, stopButton)

	mainWindow.SetContent(container.NewCenter(cont))
	mainWindow.Show()
	App.Run()
}

func plusOrMinus(mainCount *int, count int, infoPanel *widget.Label) {
	*mainCount += count
	infoPanel.SetText(strconv.Itoa(*mainCount))
}

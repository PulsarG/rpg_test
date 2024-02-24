package main

import (
	//"context"
	//"log"

	//"fmt"
	"strconv"
	"time"

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

	textLabel := widget.NewLabel("EDGE")
	mainCount := 0
	infoPanel := widget.NewLabel(strconv.Itoa(mainCount))

	stopButton := widget.NewButton("Минус", func() { plusOrMinus(&mainCount, -1, infoPanel) })

	startButton := widget.NewButton("Плюс", func() { plusOrMinus(&mainCount, 1, infoPanel) })

	cont := container.NewVBox(startButton, stopButton)
	firstContainer := container.NewHBox(textLabel, infoPanel, cont)

	pausCount := 0
	pausLabel := widget.NewLabel(strconv.Itoa(pausCount))
	pausCont := container.NewHBox(pausLabel)
	go toPausCount(&pausCount, pausLabel)

	mainCont := container.NewVBox(firstContainer, pausCont)
	mainWindow.SetContent(mainCont)
	mainWindow.Show()
	App.Run()
}

func toPausCount(count *int, label *widget.Label) {
	tiker := time.NewTicker(1 * time.Second)
	for range tiker.C {
		*count += 1
		label.SetText(strconv.Itoa(*count))
	}
}

func plusOrMinus(mainCount *int, count int, infoPanel *widget.Label) {
	*mainCount += count
	infoPanel.SetText(strconv.Itoa(*mainCount))
}

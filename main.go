// 0.04

package main

import (
	//"context"
	//"log"

	"fmt"
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
	go toPausCount(&pausCount, &mainCount, pausLabel)

	mainCont := container.NewVBox(firstContainer, pausCont)
	mainWindow.SetContent(mainCont)
	mainWindow.Show()
	App.Run()
}

func toPausCount(count *int, mainCount *int, label *widget.Label) {
	paus := 0
	mainPercent := 20
	pausPercent := 20
	tiker := time.NewTicker(1 * time.Second)
	fmt.Println(*mainCount, paus, mainPercent)

	for range tiker.C {
		if mainPercent == 0 {
			tiker.Stop()
		}
		if *count == (mainPercent - (10 * (*mainCount))) {
			fmt.Println(*mainCount, paus, mainPercent)
			if paus == (pausPercent - (3 * (*mainCount))) {
				*count = 0
				paus = 0
				continue
			} else {
				paus += 1
				label.SetText(strconv.Itoa(*count) + " " + strconv.Itoa(paus))
			}
		} else {
			*count += 1
			fmt.Println(*mainCount, paus, mainPercent)
			label.SetText(strconv.Itoa(*count) + " " + strconv.Itoa(paus))
		}
	}
}

func plusOrMinus(mainCount *int, count int, infoPanel *widget.Label) {
	*mainCount += count
	infoPanel.SetText(strconv.Itoa(*mainCount))
}

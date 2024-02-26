// 0.05

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

	"rpg_test/data"
)

func main() {
	App := app.New()
	mainWindow := App.NewWindow("RPG TEST")
	mainWindow.Resize(fyne.NewSize(500, 500))

	char := data.GetStartChar()

	textLabel := widget.NewLabel("EDGE")
	//mainCount := 0
	infoPanel := widget.NewLabel(strconv.Itoa(char.GetMainCount()))

	stopButton := widget.NewButton("Минус", func() { plusOrMinus(char, -1, infoPanel) })
	startButton := widget.NewButton("Плюс", func() { plusOrMinus(char, 1, infoPanel) })

	cont := container.NewVBox(startButton, stopButton)
	firstContainer := container.NewHBox(textLabel, infoPanel, cont)

	//pausCount := 0
	pausLabel := widget.NewLabel(strconv.Itoa(char.GetPausCount()))
	pausCont := container.NewHBox(pausLabel)
	go toPausCount(char, pausLabel)

	mainCont := container.NewVBox(firstContainer, pausCont)
	mainWindow.SetContent(mainCont)
	mainWindow.Show()
	App.Run()
}

func toPausCount(char *data.DataChar, label *widget.Label) {
	paus := 0
	//mainPercent := 60
	//pausPercent := 20
	tiker := time.NewTicker(1 * time.Second)
	//fmt.Println(*mainCount, paus, mainPercent)

	for range tiker.C {
		if char.GetMainPercent() == 0 {
			tiker.Stop()
		}
		if char.GetPausCount() == (char.GetMainPercent() - (10 * (char.GetMainCount()))) {
			fmt.Println(char.GetMainCount(), paus, char.GetMainPercent())
			if paus == (char.GetPausPercent() - (3 * (char.GetMainCount()))) {
				char.SetPausCount(0)
				paus = 0
				continue
			} else {
				paus += 1
				label.SetText(strconv.Itoa(char.GetPausCount()) + " " + strconv.Itoa(paus))
			}
		} else {
			char.SetPausCount(char.GetPausCount() + 1)
			fmt.Println(char.GetMainCount(), paus, char.GetMainPercent())
			label.SetText(strconv.Itoa(char.GetPausCount()) + " " + strconv.Itoa(paus))
		}
	}
}

func plusOrMinus(char *data.DataChar, count int, infoPanel *widget.Label) {
	char.SetMainCount(count)
	infoPanel.SetText(strconv.Itoa(char.GetMainPercent()))
}

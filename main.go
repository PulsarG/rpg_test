// ** 0.07

package main

import (
	//"context"
	//"log"

	//"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"github.com/gen2brain/beeep"
	"strconv"
	"time"

	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/dialog"
	//"fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	d "rpg_test/data"
	"rpg_test/notyfication"
)

func main() {
	App := app.New()
	mainWindow := App.NewWindow(d.NameWindow)
	mainWindow.Resize(fyne.NewSize(d.W_Window, d.H_Window))

	char := d.GetStartChar()

	textLabel := widget.NewLabel("EDGE")

	infoPanel := widget.NewLabel(strconv.Itoa(char.GetMainCount()))

	stopButton := widget.NewButton("-", func() { plusOrMinus(char, -1, infoPanel) })
	startButton := widget.NewButton("+", func() { plusOrMinus(char, 1, infoPanel) })

	cont := container.NewVBox(startButton, stopButton)
	firstContainer := container.NewHBox(textLabel, infoPanel, cont)

	pausLabel := widget.NewLabel(strconv.Itoa(char.GetPausCount()))
	pausCont := container.NewHBox(pausLabel)
	go toPausCount(char, pausLabel)

	mainCont := container.NewVBox(firstContainer, pausCont)
	mainWindow.SetContent(mainCont)
	mainWindow.Show()
	App.Run()
}

// !!! Брать значения из структуры только для создания внутренней переменной
// !!! с которой и работать. Вносить изменения в поля структуры только при
// !!! выходе из программы

func toPausCount(char *d.DataChar, label *widget.Label) {
	paus := 0
	toPaus := 0
	//mainPercent := 60
	//pausPercent := 20
	tiker := time.NewTicker(1 * time.Second)

	for range tiker.C {
		if char.GetMainPercent() <= 0 {
			tiker.Stop()
			label.SetText(d.LooseText)
			return
		}

		if char.GetPausCount() >= char.GetMainPercent() {
			// * Пауза
			// *  Начало паузы
			if paus == 0 {
				notyfication.Beep("test on")
			}
			paus += 1
			label.SetText(strconv.Itoa(char.GetPausCount()) + " " + strconv.Itoa(paus))
			// * Конец паузы
			if (char.GetPausPercent() - paus) == 0 {
				paus = 0
				notyfication.Beep("test off")
				char.SetPausCount(0)
				continue
			}

		} else {
			// * Основной счетчик
			toPaus += 1
			char.SetPausCount(char.GetPausCount() + 1)
			label.SetText(strconv.Itoa(char.GetPausCount()) + " " + strconv.Itoa(paus))
		}
	}
}

func plusOrMinus(char *d.DataChar, count int, infoPanel *widget.Label) {
	char.SetMainCount(count)
	char.SetMainPercent((char.GetMainPercent() - (count * d.Step)))
	infoPanel.SetText(strconv.Itoa(char.GetMainCount()))
}

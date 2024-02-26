// ** 0.08

package main

import (
	//"context"
	//"log"

	//"fmt"
	"fmt"

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
	hardCountLabel := widget.NewLabel(strconv.Itoa(char.GetHardCount()))

	stopButton := widget.NewButton("-", func() { plusOrMinus(char, -1, infoPanel) })
	startButton := widget.NewButton("+", func() { plusOrMinus(char, 1, infoPanel) })

	cont := container.NewVBox(startButton, stopButton)
	firstContainer := container.NewHBox(textLabel, infoPanel, cont, hardCountLabel)

	pausLabel := widget.NewLabel(strconv.Itoa(char.GetPausCount()))
	pausCont := container.NewHBox(pausLabel)
	go toPausCount(char, pausLabel, hardCountLabel)

	mainCont := container.NewVBox(firstContainer, pausCont)
	mainWindow.SetContent(mainCont)
	mainWindow.Show()
	App.Run()
}

// !!! Брать значения из структуры только для создания внутренней переменной
// !!! с которой и работать. Вносить изменения в поля структуры только при
// !!! выходе из программы

// !!! Разбить на методы
func toPausCount(char *d.DataChar, label, hardLabel *widget.Label) {
	paus := 0
	toPaus := 0
	hardStuck := 0
	tiker := time.NewTicker(1 * time.Second)

	for range tiker.C {
		if char.GetMainPercent() <= 0 {
			tiker.Stop()
			label.SetText(d.LooseText)
			return
		}

		if toPaus == char.GetMainPercent() {
			// * Пауза
			// *  Начало паузы
			fmt.Println(paus)
			if paus == 0 {
				char.SetHardCount(char.GetHardCount() + 1)
				if (char.GetHardCount() % 5) == 0 {
					hardStuck++
					hardLabel.SetText(strconv.Itoa(char.GetHardCount()) + " " + strconv.Itoa(hardStuck))
					char.SetMainPercent(char.GetMainPercent() + 3)
					char.SetPausPercent(char.GetPausPercent() + 10)
				}
				hardLabel.SetText(strconv.Itoa(char.GetHardCount()) + " " + strconv.Itoa(hardStuck))
				notyfication.Beep("test on")
			}
			paus += 1
			label.SetText(strconv.Itoa(char.GetPausCount()) + " пауза " + strconv.Itoa(paus))
			// * Конец паузы
			if char.GetPausPercent() == paus {
				fmt.Println(char.GetPausPercent(), paus)
				paus = 0
				notyfication.Beep("test off")
				toPaus = 0
				continue
			}

		} else {
			fmt.Println(char.GetPausPercent(), paus)
			// * Основной счетчик
			toPaus += 1
			//char.SetPausCount(char.GetPausCount() + 1)
			label.SetText(strconv.Itoa(toPaus) + " * " + strconv.Itoa(paus))
		}
	}
}

func plusOrMinus(char *d.DataChar, count int, infoPanel *widget.Label) {
	char.SetMainCount(count)
	char.SetMainPercent((char.GetMainPercent() - (count * d.Step)))
	char.SetPausPercent((char.GetPausPercent() - (count * d.PausStep)))
	infoPanel.SetText(strconv.Itoa(char.GetMainCount()))
}

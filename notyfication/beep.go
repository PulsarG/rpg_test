package notyfication

import "github.com/gen2brain/beeep"

func Beep(title string) {
	err := beeep.Notify(title, title, "assets/information.png")
	if err != nil {
		panic(err)
	}
	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}

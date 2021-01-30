package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		buttonSHA256 := ui.NewButton("Encrypt to SHA256")
		buttonMD5 := ui.NewButton("Encrypt to MD5")
		crypted := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter a string to encrypt:"), false)
		box.Append(input, false)
		box.Append(buttonSHA256, false)
		box.Append(buttonMD5, false)
		box.Append(crypted, false)
		window := ui.NewWindow("Hello", 800, 600, false)
		window.SetMargined(true)
		window.SetChild(box)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

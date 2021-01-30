package main

import (
	"crypto-client/controllers"
	"github.com/andlabs/ui"
)

func main() {
	var encryptedS string
	err := ui.Main(func() {
		input := ui.NewEntry()
		buttonSHA256 := ui.NewButton("Encrypt to SHA256")
		buttonMD5 := ui.NewButton("Encrypt to MD5")
		crypted := ui.NewLabel("Encrypted: ")
		buttonSave := ui.NewButton("Save to file")
		sep := ui.NewLabel("")
		sep2 := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter a string to encrypt:"), false)
		box.Append(input, false)
		box.Append(buttonSHA256, false)
		box.Append(buttonMD5, false)
		box.Append(sep, false)
		box.Append(crypted, false)
		box.Append(sep2, false)
		box.Append(buttonSave, false)

		window := ui.NewWindow("Hello", 800, 600, false)
		window.SetMargined(true)
		window.SetChild(box)
		buttonSHA256.OnClicked(func(b *ui.Button) {
			str := input.Text()
			text := controllers.SHA256(str)
			crypted.SetText("Encrypted: " + text)
			encryptedS = text
		})
		buttonMD5.OnClicked(func(b *ui.Button) {
			str := input.Text()
			text := controllers.MD5(str)
			crypted.SetText("Encrypted: " + text)
			encryptedS = text
		})
		buttonSave.OnClicked(func(b *ui.Button) {
			controllers.Save(encryptedS)
		})
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

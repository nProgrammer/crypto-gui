package main

import (
	"crypto-client/views"
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		box := views.MainView()
		buttonCrypto := ui.NewButton("Go to crypto secure+")
		box.Append(buttonCrypto, false)

		window := ui.NewWindow("Hello", 800, 600, false)
		window.SetMargined(true)
		window.SetChild(box)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
		buttonCrypto.OnClicked(func(b *ui.Button) {
			box = views.CryptoView()
			window.SetChild(box)
			fmt.Println("10")
		})
	})

	if err != nil {
		panic(err)
	}
}

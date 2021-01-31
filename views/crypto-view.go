package views

import (
	"github.com/andlabs/ui"
)

func CryptoView() *ui.Box {
	sep := ui.NewLabel("Hello world!")
	box := ui.NewVerticalBox()

	box.Append(sep, false)

	return box
}

// ** 0.01

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)


func main() {
	App := app.New()
	mainWindow := App.NewWindow("Barony Trainer")
	mainWindow.Resize(fyne.NewSize(500, 500))

	mainWindow.Show()
	App.Run()
}

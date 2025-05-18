package main

import (
	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	window := a.NewWindow("Calculator")
	window.Resize(fyne.NewSize(600, 500))

	label1 := widget.NewLabel("")
	label1.Resize(fyne.NewSize(600, 100))

	//window.SetContent(Content)
	window.ShowAndRun()
}

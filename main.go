package main

import (
	"fmt"
	//"fyne.io/fyne/v2"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	window := a.NewWindow("Calculator")
	window.Resize(fyne.NewSize(500, 600))

	label1 := widget.NewLabel("")
	label1.Resize(fyne.NewSize(500, 100))

	button1 := widget.NewButton("ONE", func() {
		label1.SetText(fmt.Sprintf("1"))
	})
	button1.Resize(fyne.NewSize(100, 100))

	content := container.NewVBox(label1, button1)
	window.SetContent(content)
	window.ShowAndRun()
}

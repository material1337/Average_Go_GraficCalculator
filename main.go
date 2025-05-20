package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/expr-lang/expr"
	"image/color"
)

func main() {
	//App
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	//Display
	display := widget.NewEntry()
	display.SetPlaceHolder("0")
	display.Disable()

	display.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	displayContainer := container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.White), display)

	//For monitoring the status of brackets
	isOpenBracket := true

	//Add symbol to label(Entry)
	appendToDisplay := func(char string) {
		display.SetText(display.Text + char)
	}

	//Full clear func
	clearDisplay := func() {
		display.SetText("")
		isOpenBracket = true
	}

	//Brackets func
	toggleBracket := func() {
		if isOpenBracket {
			appendToDisplay("(")
		} else {
			appendToDisplay(")")
		}
		isOpenBracket = !isOpenBracket //Switch
	}

	//Delete last char  func( very hard to write)
	deleteLastChar := func() {
		currentText := display.Text
		if len(currentText) > 0 {

			newText := currentText[:len(currentText)-1]
			display.SetText(newText)

			if len(newText) > 0 {
				lastChar := string(newText[len(newText)-1])
				if lastChar == "(" {
					isOpenBracket = false
				} else if lastChar == ")" {
					isOpenBracket = true
				}
			} else {
				isOpenBracket = true
			}
		}
	}

	//Main calculation func
	calculate := func() {
		expression := display.Text
		if expression == "" {
			return
		}
		result, err := expr.Eval(expression, nil)
		if err != nil {
			display.SetText("Error")
			return
		}
		display.SetText(fmt.Sprintf("%v", result))
		isOpenBracket = true
	}

	//Buttons
	button1 := widget.NewButton("1", func() { appendToDisplay("1") })
	button2 := widget.NewButton("2", func() { appendToDisplay("2") })
	button3 := widget.NewButton("3", func() { appendToDisplay("3") })
	button4 := widget.NewButton("4", func() { appendToDisplay("4") })
	button5 := widget.NewButton("5", func() { appendToDisplay("5") })
	button6 := widget.NewButton("6", func() { appendToDisplay("6") })
	button7 := widget.NewButton("7", func() { appendToDisplay("7") })
	button8 := widget.NewButton("8", func() { appendToDisplay("8") })
	button9 := widget.NewButton("9", func() { appendToDisplay("9") })
	button10 := widget.NewButton("AC", func() { clearDisplay() })
	button11 := widget.NewButton("()", func() { toggleBracket() })
	button12 := widget.NewButton("%", func() { appendToDisplay("%") })
	button13 := widget.NewButton("/", func() { appendToDisplay("/") })
	button14 := widget.NewButton("*", func() { appendToDisplay("*") })
	button15 := widget.NewButton("-", func() { appendToDisplay("-") })
	button16 := widget.NewButton("+", func() { appendToDisplay("+") })
	button17 := widget.NewButton("=", func() { calculate() })
	button18 := widget.NewButton("Del", func() { deleteLastChar() })
	button19 := widget.NewButton(",", func() { appendToDisplay(",") })
	button20 := widget.NewButton("0", func() { appendToDisplay("0") })

	//GridLayout 4 paragraphs
	grid := container.New(layout.NewGridLayout(4),
		button10, button11, button12, button13,
		button7, button8, button9, button14,
		button4, button5, button6, button15,
		button1, button2, button3, button16,
		button20, button19, button18, button17)

	//button + entry
	content := container.NewVBox(
		displayContainer,
		grid,
	)
	//End
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 250))
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}

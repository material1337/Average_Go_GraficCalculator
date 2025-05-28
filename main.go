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
	"regexp"
	"strings"
)

func main() {
	// App
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	// Display
	display := widget.NewEntry()
	display.SetPlaceHolder("0")
	display.Disable()

	display.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	displayContainer := container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.White), display)

	// For monitoring the status of brackets
	isOpenBracket := true

	// Add symbol to label (Entry)
	appendToDisplay := func(char string) {
		display.SetText(display.Text + char)
	}

	// Full clear func
	clearDisplay := func() {
		display.SetText("")
		isOpenBracket = true
	}

	// Brackets func
	toggleBracket := func() {
		if isOpenBracket {
			appendToDisplay("(")
		} else {
			appendToDisplay(")")
		}
		isOpenBracket = !isOpenBracket // Switch
	}

	// Percent func
	percent := func() {
		currentText := display.Text
		if currentText == "" {
			return
		}

		// Find the last number before the percent sign
		re := regexp.MustCompile(`(\d+\.?\d*)\s*$`)
		matches := re.FindStringSubmatch(currentText)
		if len(matches) < 2 {
			return // No valid number found
		}

		// Append the percentage calculation
		number := matches[1]
		// Transform e.g., "50" to "(50 / 100)"
		percentageExpr := fmt.Sprintf("(%s / 100)", number)
		// Replace the number with the percentage expression
		newText := re.ReplaceAllString(currentText, percentageExpr)
		display.SetText(newText)
	}

	// Delete last char func
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

	// Main calculation func
	calculate := func() {
		expression := display.Text
		if expression == "" {
			return
		}

		// Replace commas with dots for decimal numbers
		expression = strings.ReplaceAll(expression, ",", ".")

		// Handle percentage expressions
		// e.g., "50 + 10%" becomes "50 + (10 / 100 * 50)"
		re := regexp.MustCompile(`(\d+\.?\d*)\s*%`)
		for re.MatchString(expression) {
			matches := re.FindStringSubmatch(expression)
			if len(matches) < 2 {
				break
			}
			number := matches[1]
			// Find the base number before the percentage
			// We look for a number or expression before the percentage
			baseRe := regexp.MustCompile(`(\d+\.?\d*)\s*[\+\-\*/]\s*\d+\.?\d*\s*%`)
			baseMatches := baseRe.FindStringSubmatch(expression)
			var base string
			if len(baseMatches) > 1 {
				base = baseMatches[1]
			} else {
				base = "1" // Default base for cases like "10%"
			}
			percentageExpr := fmt.Sprintf("(%s / 100 * %s)", number, base)
			expression = re.ReplaceAllString(expression, percentageExpr)
		}

		result, err := expr.Eval(expression, nil)
		if err != nil {
			display.SetText("Error")
			return
		}
		display.SetText(fmt.Sprintf("%v", result))
		isOpenBracket = true
	}

	// Buttons
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
	button12 := widget.NewButton("%", func() { percent() })
	button13 := widget.NewButton("/", func() { appendToDisplay("/") })
	button14 := widget.NewButton("*", func() { appendToDisplay("*") })
	button15 := widget.NewButton("-", func() { appendToDisplay("-") })
	button16 := widget.NewButton("+", func() { appendToDisplay("+") })
	button17 := widget.NewButton("=", func() { calculate() })
	button18 := widget.NewButton("Del", func() { deleteLastChar() })
	button19 := widget.NewButton(",", func() { appendToDisplay(".") })
	button20 := widget.NewButton("0", func() { appendToDisplay("0") })

	// GridLayout 4 columns
	grid := container.New(layout.NewGridLayout(4),
		button10, button11, button12, button13,
		button7, button8, button9, button14,
		button4, button5, button6, button15,
		button1, button2, button3, button16,
		button20, button19, button18, button17)

	// Button + entry
	content := container.NewVBox(
		displayContainer,
		grid,
	)
	// End
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 250))
	myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}

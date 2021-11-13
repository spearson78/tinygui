package calc

import (
	"github.com/shopspring/decimal"
	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/component"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/layout"
)

var btnSeven component.ButtonState
var btnEignt component.ButtonState
var btnNine component.ButtonState

var btnFour component.ButtonState
var btnFive component.ButtonState
var btnSix component.ButtonState

var btnOne component.ButtonState
var btnTwo component.ButtonState
var btnThree component.ButtonState

var btnZero component.ButtonState
var btnDot component.ButtonState
var btnNegate component.ButtonState

var btnDiv component.ButtonState
var btnMul component.ButtonState
var btnSub component.ButtonState
var btnAdd component.ButtonState
var btnEquals component.ButtonState

var btnClear component.ButtonState
var btnDel component.ButtonState

var txtResultDisplay component.TextFieldState

var calcShift = int32(1)
var valShift = int32(0)
var accumulator = decimal.Decimal{}
var calcValue = decimal.Decimal{}

type Operator byte

const (
	Equals Operator = iota
	Div
	Mul
	Sub
	Add
)

var op Operator

func applyLastOperator(newOp Operator) {
	switch op {
	case Equals:
		accumulator = calcValue
	case Div:
		accumulator = accumulator.Div(calcValue)
	case Mul:
		accumulator = accumulator.Mul(calcValue)
	case Sub:
		accumulator = accumulator.Sub(calcValue)
	case Add:
		accumulator = accumulator.Add(calcValue)
	}
	calcValue = decimal.Decimal{}
	calcShift = int32(1)
	valShift = int32(0)
	op = newOp
}

func doNumericButton(val int32) {
	calcValue = calcValue.Shift(calcShift)
	calcValue = calcValue.Add(decimal.NewFromInt32(val).Shift(valShift))
	if valShift != 0 {
		valShift--
	}
}

//go:noinline
func calcRow1(g *tinygui.GuiContext, btnGrid *layout.GridLayout, size component.ComponentSize) {

	btnProps := component.NewButtonProps(g)
	btnProps.ComponentSize = size
	btnProps.Style = component.Text

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "7"
	if component.Button(g, &btnSeven, &btnProps) {
		doNumericButton(7)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "8"
	if component.Button(g, &btnEignt, &btnProps) {
		doNumericButton(8)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "9"
	if component.Button(g, &btnNine, &btnProps) {
		doNumericButton(9)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "/"
	btnProps.Color = g.Theme.Text
	if component.Button(g, &btnDiv, &btnProps) {
		applyLastOperator(Div)
	}

	btnProps.ComponentPos = btnGrid.EndRow()
	btnProps.Style = component.Contained
	btnProps.Label = "C"
	btnProps.Color = g.Theme.ErrorColor
	btnProps.DisableElevation = true
	if component.Button(g, &btnClear, &btnProps) {
		op = Equals
		accumulator = decimal.Decimal{}
		calcValue = decimal.Decimal{}
		calcShift = int32(1)
		valShift = int32(0)
	}
}

//go:noinline
func calcRow2(g *tinygui.GuiContext, btnGrid *layout.GridLayout, size component.ComponentSize) {

	btnProps := component.NewButtonProps(g)
	btnProps.ComponentSize = size
	btnProps.Style = component.Text

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "4"
	if component.Button(g, &btnFour, &btnProps) {
		doNumericButton(4)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "5"
	if component.Button(g, &btnFive, &btnProps) {
		doNumericButton(5)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "6"
	if component.Button(g, &btnSix, &btnProps) {
		doNumericButton(6)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "X"
	btnProps.Color = g.Theme.Text
	if component.Button(g, &btnMul, &btnProps) {
		applyLastOperator(Mul)
	}

	btnProps.ComponentPos = btnGrid.EndRow()
	btnProps.Label = ""
	btnProps.Color = g.Theme.ErrorColor
	btnProps.Style = component.Contained
	btnProps.StartIcon = icon.Backspace
	btnProps.DisableElevation = true
	if component.Button(g, &btnDel, &btnProps) {
		if valShift != 0 {
			valShift++
			calcValue = calcValue.Truncate(-(valShift + 1))
			if valShift == -1 {
				calcShift = int32(1)
				valShift = int32(0)
			}
		} else {
			calcValue = calcValue.Shift(-1).Truncate(0)
		}
	}
}

//go:noinline
func calcRow3(g *tinygui.GuiContext, btnGrid *layout.GridLayout, size component.ComponentSize) {

	btnProps := component.NewButtonProps(g)
	btnProps.ComponentSize = size
	btnProps.Style = component.Text

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "1"
	if component.Button(g, &btnOne, &btnProps) {
		doNumericButton(1)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "2"
	if component.Button(g, &btnTwo, &btnProps) {
		doNumericButton(2)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "3"
	if component.Button(g, &btnThree, &btnProps) {
		doNumericButton(3)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "-"
	btnProps.Color = g.Theme.Text
	if component.Button(g, &btnSub, &btnProps) {
		applyLastOperator(Sub)
	}

	btnProps.ComponentPos = btnGrid.EndRow()
	btnProps.ComponentSize.H = btnProps.ComponentSize.H * 2
	btnProps.Label = "="
	btnProps.Color = g.Theme.DefaultColor
	btnProps.Style = component.Contained
	btnProps.DisableElevation = true
	if component.Button(g, &btnEquals, &btnProps) {
		applyLastOperator(Equals)
	}
}

//go:noinline
func calcRow4(g *tinygui.GuiContext, btnGrid *layout.GridLayout, size component.ComponentSize) {

	btnProps := component.NewButtonProps(g)
	btnProps.ComponentSize = size
	btnProps.Style = component.Text

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "+/-"
	btnProps.Color = g.Theme.Text
	if component.Button(g, &btnNegate, &btnProps) {
		if calcValue.Equal(decimal.Zero) {
			accumulator = accumulator.Neg()
		} else {
			calcValue = calcValue.Neg()
		}
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "0"
	btnProps.Color = g.Theme.DefaultColor
	if component.Button(g, &btnZero, &btnProps) {
		doNumericButton(0)
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "."
	btnProps.Color = g.Theme.Text
	if component.Button(g, &btnDot, &btnProps) {
		if valShift == 0 {
			calcShift = 0
			valShift = -1
		}
	}

	btnProps.ComponentPos = btnGrid.NextCell()
	btnProps.Label = "+"
	if component.Button(g, &btnAdd, &btnProps) {
		applyLastOperator(Add)
	}
}

//go:noinline
func calcDisplay(g *tinygui.GuiContext) {
	if calcValue.Equal(decimal.Zero) {
		txtResultDisplay.Text = accumulator.String()
	} else {
		txtResultDisplay.Text = calcValue.String()
	}

	txtProps := component.NewTextFieldProps(g)
	txtProps.X = 5
	txtProps.Y = 5
	txtProps.W = 230
	component.TextField(g, &txtResultDisplay, &txtProps)
}

//go:noinline
func CalcGui(g *tinygui.GuiContext) {

	btnSize := component.ComponentSize{
		W: 46,
		H: 61,
	}

	btnGrid := layout.Grid(0, 35, btnSize.W, btnSize.H, 1)

	calcRow1(g, &btnGrid, btnSize)
	calcRow2(g, &btnGrid, btnSize)
	calcRow3(g, &btnGrid, btnSize)
	calcRow4(g, &btnGrid, btnSize)
	calcDisplay(g)
}

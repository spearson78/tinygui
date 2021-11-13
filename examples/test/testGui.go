package test

import (
	"strconv"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/component"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/layout"
)

var btn1 component.ButtonState
var btn2 component.ButtonState
var btn3 component.ButtonState
var btn4 component.IconButtonState
var ledOn component.CheckBoxState
var fab1State component.FloatingActionButtonState
var sliderState = component.SliderState{Value: 0.5}
var sliderText = component.TextFieldState{Text: "0"}
var txtCounterState = component.DashCardState{
	Value:       "0",
	ActionIcon:  icon.Add,
	ActionLabel: "Reset",
}
var counter = int64(0)
var switchState = component.SwitchState{}

/*
var fab2State tinygui.FloatingActionButtonState
var radio1State tinygui.RadioButtonState
var radio2State tinygui.RadioButtonState
var radio3State tinygui.RadioButtonState

*/
//go:noinline
func TestGui(g *tinygui.GuiContext) {

	btnProps := component.NewButtonProps(g)
	btnProps.ComponentSize = component.ComponentSize{
		W: 110,
		H: 30,
	}

	grid := layout.Grid(5, 5, btnProps.ComponentSize.W, btnProps.ComponentSize.H+5, 2)

	btnProps.ComponentPos = grid.NextCell()
	btnProps.Label = "ON"
	btnProps.Style = component.Contained
	btnProps.StartIcon = icon.Add
	if component.Button(g, &btn1, &btnProps) {
		ledOn.Checked = true
		g.UpdateState()
	}

	icnProps := component.NewIconButtonProps(g)
	icnProps.ComponentPos = grid.EndRow()
	icnProps.Icon = icon.Edit
	if component.IconButton(g, &btn4, &icnProps) {
		ledOn.Checked = false
		g.UpdateState()
	}

	btnProps.ComponentPos = grid.NextCell()
	btnProps.Label = "Off"
	btnProps.Style = component.Outlined
	btnProps.StartIcon = nil
	btnProps.Color = g.Theme.SecondaryColor
	if component.Button(g, &btn3, &btnProps) {
		ledOn.Checked = false
		g.UpdateState()
	}

	chkBoxProps := component.NewCheckBoxProps(g)
	chkBoxProps.ComponentPos = grid.EndRow()
	chkBoxProps.Icon = icon.Checkmark
	chkBoxProps.Label = "TOGGLE"
	if component.CheckBox(g, &ledOn, &chkBoxProps) {
	}

	btnProps.ComponentPos = grid.NextCell()
	btnProps.Label = "ON2"
	btnProps.Style = component.Text
	btnProps.StartIcon = nil
	btnProps.EndIcon = icon.Edit
	btnProps.Color = g.Theme.SecondaryColor
	if component.Button(g, &btn2, &btnProps) {
		ledOn.Checked = true
		g.UpdateState()
	}

	txtProps := component.NewTextFieldProps(g)
	txtProps.ComponentPos = grid.EndRow()
	txtProps.ComponentSize = btnProps.ComponentSize
	if component.TextField(g, &sliderText, &txtProps) {
	}

	sldProps := component.NewSliderProps(g)
	sldProps.ComponentPos = grid.EndRow()
	sldProps.ComponentSize.W = 220
	if component.Slider(g, &sliderState, &sldProps) {
		sliderText.Text = strconv.FormatFloat(float64(sliderState.Value), 'f', 2, 64)
	}

	swtchProps := component.NewSwitchProps(g)
	swtchProps.ComponentPos = grid.EndRow()
	if component.Switch(g, &switchState, &swtchProps) {
	}

	dshProps := component.NewDashCardProps(g)
	dshProps.ComponentPos = grid.EndRow()
	dshProps.Color = g.Theme.SecondaryColor
	dshProps.Label = "Count"

	if component.DashCard(g, &txtCounterState, &dshProps) {
		counter = 0
		tempStr := strconv.FormatInt(counter, 10)
		txtCounterState.Value = tempStr
		g.UpdateState()
	}

	fltBtnProps := component.NewFloatingActionButtonProps(g)
	fltBtnProps.X = 195
	fltBtnProps.Y = 235
	fltBtnProps.Icon = icon.Edit
	if component.FloatingActionButton(g, &fab1State, &fltBtnProps) {
		counter++
		tempStr := strconv.FormatInt(counter, 10)
		txtCounterState.Value = tempStr
		g.UpdateState()
	}

	/*





		if tinygui.RadioButton(g, &radio1State,
			tinygui.Pos(10, 105),
			tinygui.Label("RADIO 1"),
			tinygui.RadioButtonCheckedState(1),
		) {
		}

		if tinygui.RadioButton(g, &radio2State,
			tinygui.Pos(10, 135),
			tinygui.Label("RADIO 2"),
			tinygui.Color(g.Theme.SecondaryColor),
			tinygui.RadioButtonCheckedState(2),
		) {
		}

		if tinygui.RadioButton(g, &radio3State,
			tinygui.Pos(10, 165),
			tinygui.Label("RADIO 3"),
			tinygui.Disabled(true),
			tinygui.RadioButtonCheckedState(3),
		) {
		}



		if tinygui.FloatingActionButton(g, &fab2State,
			tinygui.Pos(150, 100),
			tinygui.Color(g.Theme.SecondaryColor),
			tinygui.Icon(tinygui.Edit),
		) {
			ledOn.Checked = !ledOn.Checked
			g.UpdateState()
		}

		tinygui.DashCard(g, &txtTempDisplayState,
			tinygui.Pos(10, 230),
		)

		//machine.LED.Set(ledOn.Checked)
	*/

}

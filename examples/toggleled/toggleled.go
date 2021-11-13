package toggleled

import (
	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/component"

	"machine"
)

//State variable to track whether LED is on or off
var LedState bool

//State for Toggle Button
var btnToggleLed component.ButtonState

//Gui Function GuiContext is provided by Tiny Gui
func ToggleLedGui(g *tinygui.GuiContext) {

	//Create button props including default values based on current theme
	btnToggleProps := component.NewButtonProps(g)

	//Set position and Size of the button
	btnToggleProps.X = 10
	btnToggleProps.Y = 10
	btnToggleProps.W = 220
	btnToggleProps.H = 50

	//Set the label of the button
	btnToggleProps.Label = "TOGGLE LED"

	//Set the style of the button
	btnToggleProps.Style = component.Contained

	//Render the Button and handle the click
	if component.Button(g, &btnToggleLed, &btnToggleProps) {
		//component.Button returns true if the button was clicked
		LedState = !LedState
		machine.LED.Set(LedState)
	}
}

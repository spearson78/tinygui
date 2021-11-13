package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/primitive"
)

type TextFieldStyle byte

const (
	Standard TextFieldStyle = iota
	Filled
	OutlinedTextField
)

type TextFieldState struct {
	Text string
}

type TextFieldProps struct {
	ComponentPos
	ComponentSize
	Style    TextFieldStyle
	Disabled bool
	Color    color.RGBA
}

func NewTextFieldProps(g *tinygui.GuiContext) TextFieldProps {
	return TextFieldProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: 100,
			H: -1,
		},
		Style:    Standard,
		Disabled: false,
		Color:    g.Theme.Text,
	}
}

//go:noinline
func TextField(g *tinygui.GuiContext, state *TextFieldState, props *TextFieldProps) bool {

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, props.W, 30)
	if eventAction == Ignore || eventAction == Click {
		return false
	}

	if (g.Event.Type & event.Invalidate) != 0 {
		g.Display.FillRect(props.X, props.Y, props.W, 30, g.Theme.Background)
		g.Display.FillRect(props.X, props.Y+28, props.W, 1, props.Color)
	} else {
		g.Display.FillRect(props.X, props.Y, props.W, 27, g.Theme.Background)
	}

	primitive.WriteLine(g.Display, g.Theme.Font, props.X, int16(props.Y+21), state.Text, props.Color)

	return false
}

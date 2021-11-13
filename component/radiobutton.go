package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type RadioButtonGroup uint8

var DefaultRadioButtonGroup RadioButtonGroup

type RadioButtonProps struct {
	X            int16
	Y            int16
	Label        string
	Color        color.RGBA
	Disabled     bool
	Group        *RadioButtonGroup
	CheckedState uint8
}

type RadioButtonState struct {
	animateState int8
}

func NewRadioButtonProps(g *tinygui.GuiContext) RadioButtonProps {
	return RadioButtonProps{
		X:            10,
		Y:            10,
		Label:        "",
		Color:        g.Theme.DefaultColor,
		Disabled:     false,
		Group:        &DefaultRadioButtonGroup,
		CheckedState: 1,
	}
}

//go:noinline
func RadioButton(g *tinygui.GuiContext, state *RadioButtonState, props *RadioButtonProps) bool {

	clicked := false
	colBackground := g.Theme.Background
	boxColor := g.Theme.Border

	_, ow := primitive.LineWidth(g.Theme.Font, props.Label)
	w := int16(ow)
	h := int16(20)

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, w, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.animateState = 10
		*props.Group = RadioButtonGroup(props.CheckedState)
		clicked = true
		g.UpdateState()
	}

	checked := *props.Group == RadioButtonGroup(props.CheckedState)

	if props.Disabled {
		boxColor = g.Theme.DisabedColor
	} else if checked {
		boxColor = props.Color
	}

	//Clear out old widget
	if (g.Event.Type & event.Invalidate) != 0 {
		tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, int16(w), int16(h), colBackground)
	}

	if state.animateState > 0 {
		fillColor := g.Theme.Background
		fillColor = primitive.Desaturate(props.Color, 76)
		fillColor = primitive.TransitionColor(fillColor, g.Theme.Background, uint8(state.animateState*25))

		tinydraw.FilledCircleEx(g.Display, props.X+10, props.Y+10, 15, fillColor)
		state.animateState--
		g.InvalidateRect(props.X-5, props.Y-5, props.X+25, props.Y+25, false)
	}

	//Radio Button
	tinydraw.FilledCircleEx(g.Display, props.X+10, props.Y+10, 10, boxColor)
	tinydraw.FilledCircleEx(g.Display, props.X+10, props.Y+10, 8, colBackground)

	if checked {
		tinydraw.FilledCircleEx(g.Display, props.X+10, props.Y+10, 5, boxColor)
	}

	if (g.Event.Type & event.Invalidate) != 0 {
		colText := g.Theme.Text
		if props.Disabled {
			colText = boxColor
		}

		//Label
		primitive.WriteLine(g.Display, g.Theme.Font, int16(props.X+24), int16(props.Y+17), props.Label, colText)
	}

	return clicked
}

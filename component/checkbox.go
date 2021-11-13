package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type CheckBoxProps struct {
	ComponentPos
	Label    string
	Color    color.RGBA
	Disabled bool
	Icon     icon.Icon
}

type CheckBoxState struct {
	animateState int8
	Checked      bool
}

func NewCheckBoxProps(g *tinygui.GuiContext) CheckBoxProps {
	return CheckBoxProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		Label:    "",
		Disabled: false,
		Color:    g.Theme.DefaultColor,
		Icon:     icon.Checkmark,
	}
}

//go:noinline
func CheckBox(g *tinygui.GuiContext, state *CheckBoxState, props *CheckBoxProps) bool {

	clicked := false

	_, ow := primitive.LineWidth(g.Theme.Font, props.Label)
	w := int16(ow)
	h := int16(20)

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, w, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.animateState = 10
		state.Checked = !state.Checked
		clicked = true
		g.UpdateState()
	}

	colChecked := g.Theme.Background
	colBackground := g.Theme.Background

	if (g.Event.Type & event.Invalidate) != 0 {
		//Clear out old widget
		tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, int16(w), int16(h), colBackground)
	} else {
		tinydraw.FilledRectangleEx(g.Display, props.X+1, props.Y+6, 16, 16, colBackground)
	}

	//Check box
	var boxColor = g.Theme.Border

	if props.Disabled {
		boxColor = g.Theme.DefaultColor
	} else if state.Checked {
		boxColor = props.Color
	}

	if state.animateState > 0 {
		fillColor := primitive.Desaturate(props.Color, 76)
		fillColor = primitive.TransitionColor(fillColor, g.Theme.Background, uint8(state.animateState*25))

		tinydraw.FilledCircleEx(g.Display, props.X+8, props.Y+13, 15, fillColor)
		state.animateState--
		g.InvalidateRect(props.X-8, props.Y-4, props.X+24, props.Y+28, false)
	}

	if !state.Checked {
		primitive.OutlineBox(g.Display, props.X, props.Y+5, 15+2, 15+2, boxColor)
	} else {
		primitive.FilledBox(g.Display, props.X, props.Y+5, 15+2, 15+2, boxColor)
		props.Icon(g, props.X+2, props.Y+7, 13, colChecked)
	}

	if (g.Event.Type & event.Invalidate) != 0 {
		//Label
		colText := g.Theme.Text
		if props.Disabled {
			colText = boxColor
		}

		primitive.WriteLine(g.Display, g.Theme.Font, props.X+24, props.Y+20, props.Label, colText)
	}

	return clicked
}

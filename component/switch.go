package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type SwitchProps struct {
	ComponentPos
	Label    string
	Color    color.RGBA
	Disabled bool
}

type SwitchState struct {
	animateState int8
	Checked      bool
}

func NewSwitchProps(g *tinygui.GuiContext) SwitchProps {
	return SwitchProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		Label:    "Switch",
		Disabled: false,
		Color:    g.Theme.DefaultColor,
	}
}

//go:noinline
func Switch(g *tinygui.GuiContext, state *SwitchState, props *SwitchProps) bool {

	clicked := false

	_, ow := primitive.LineWidth(g.Theme.Font, props.Label)
	w := int16(48) + int16(ow)

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, w, 30)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.Checked = !state.Checked
		if state.Checked {
			state.animateState = 10
		}
		clicked = true
		g.UpdateState()
	}

	if (g.Event.Type & event.Invalidate) != 0 {
		tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, w, int16(24), g.Theme.Background)
	} else {
		tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, 48, int16(24), g.Theme.Background)
	}

	//Check box
	var switchColor = g.Theme.Background
	var backColor = g.Theme.Border
	if state.Checked {
		switchColor = props.Color
	}

	if props.Disabled {
		backColor = primitive.Desaturate(backColor, 76)
		switchColor = primitive.Desaturate(switchColor, 76)
	}

	switchX := int16(0)
	if state.Checked {
		switchX = 20
	}

	if state.animateState > 0 {
		animColor := primitive.Desaturate(switchColor, 76)
		animColor = primitive.TransitionColor(animColor, g.Theme.Background, uint8(state.animateState*25))
		primitive.FilledCircle(g.Display, props.X+switchX-3, props.Y+5-3, 26, animColor)
		state.animateState--

		g.InvalidateRect(props.X+switchX-4, props.Y, props.X+switchX+28, props.Y+24, false)
	}

	primitive.FilledBox(g.Display, props.X+10, props.Y+5+5, 20, 10, backColor)
	primitive.FilledCircleWithShadow(g.Display, props.X+switchX, props.Y+5, 20, switchColor, g.Theme.Shadow1, g.Theme.Shadow2)

	if (g.Event.Type & event.Invalidate) != 0 {
		//Label
		colText := g.Theme.Text
		if props.Disabled {
			colText = g.Theme.DisabedColor
		}

		primitive.WriteLine(g.Display, g.Theme.Font, props.X+48, props.Y+20, props.Label, colText)
	}

	return clicked
}

package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type IconButtonProps struct {
	ComponentPos
	ComponentSize
	Disabled bool
	Color    color.RGBA
	Icon     icon.Icon
}

type IconButtonState struct {
	animateState int8
}

func NewIconButtonProps(g *tinygui.GuiContext) IconButtonProps {
	return IconButtonProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: 30,
			H: -1,
		},
		Disabled: false,
		Color:    g.Theme.DefaultColor,
		Icon:     icon.Add,
	}
}

//go:noinline
func IconButton(g *tinygui.GuiContext, state *IconButtonState, props *IconButtonProps) bool {

	w := props.W
	h := w
	centerOffset := w / 2
	clicked := false

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, true, props.X, props.Y, props.W, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.animateState = 10
		clicked = true
		g.UpdateState()
	}

	if state.animateState >= 0 {
		fillColor := g.Theme.Background
		if state.animateState > 0 {
			fillColor = primitive.Desaturate(props.Color, 76)
			fillColor = primitive.TransitionColor(fillColor, g.Theme.Background, uint8(state.animateState*25))
		}

		tinydraw.FilledCircleEx(g.Display, props.X+centerOffset, props.Y+centerOffset, (w/2)-1, fillColor)
		state.animateState--
		g.InvalidateRect(props.X, props.Y, props.X+w, props.Y+h, false)
	}

	darkColor := props.Color
	if props.Disabled {
		darkColor = g.Theme.DisabedColor
	}

	iconOffset := w / 4
	props.Icon(g, props.X+iconOffset+1, props.Y+iconOffset+1, (w / 2), darkColor)

	return clicked

}

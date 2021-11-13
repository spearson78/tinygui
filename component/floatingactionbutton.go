package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/primitive"
)

type FloatingActionButtonStyle byte

const (
	IconStyle FloatingActionButtonStyle = iota
	Extended
)

type FloatingActionButtonProps struct {
	ComponentPos
	ComponentSize
	Style    FloatingActionButtonStyle
	Disabled bool
	Color    color.RGBA
	Icon     icon.Icon
	Label    string //Ony used in Extended style
}

type FloatingActionButtonState struct {
	animateState int8
}

func NewFloatingActionButtonProps(g *tinygui.GuiContext) FloatingActionButtonProps {
	return FloatingActionButtonProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: 40,
			H: -1,
		},
		Style:    IconStyle,
		Label:    "",
		Disabled: false,
		Color:    g.Theme.DefaultColor,
		Icon:     icon.Add,
	}
}

//go:noinline
func FloatingActionButton(g *tinygui.GuiContext, state *FloatingActionButtonState, props *FloatingActionButtonProps) bool {

	w := props.W
	h := props.W
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

	darkColor := props.Color
	lightColor := g.Theme.Background
	if props.Disabled {
		darkColor = g.Theme.DisabedColor
		lightColor = color.RGBA{166, 167, 168, 255}
	}

	if state.animateState > 0 {

		fillColor := darkColor
		darkColor = primitive.Desaturate(darkColor, 76)
		darkColor = primitive.TransitionColor(darkColor, fillColor, uint8(state.animateState*25))
		state.animateState--
		g.InvalidateRect(props.X, props.Y, props.X+w, props.Y+h, false)
	}

	textColor := lightColor

	if !props.Disabled {
		shadow1 := g.Theme.Shadow1
		shadow2 := g.Theme.Shadow2

		primitive.FilledCircleWithShadow(g.Display, props.X, props.Y, props.W, darkColor, shadow1, shadow2)
	} else {
		primitive.FilledCircle(g.Display, props.X, props.Y, props.W, darkColor)
	}

	iconWidth := props.W / 2 // int16(float32(radius) * float32(0.9))
	centreOffset := (w/2 - iconWidth/2)
	if centreOffset < 0 {
		centreOffset = 4
	}

	props.Icon(g, props.X+centreOffset, props.Y+centreOffset-1, iconWidth, textColor)

	return clicked

}

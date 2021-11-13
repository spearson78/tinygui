package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type SliderState struct {
	animateState int8
	Value        float32
}

type SliderProps struct {
	ComponentPos
	ComponentSize
	Color    color.RGBA
	Disabled bool
}

func NewSliderProps(g *tinygui.GuiContext) SliderProps {
	return SliderProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: 100,
			H: -1,
		},
		Color:    g.Theme.DefaultColor,
		Disabled: false,
	}
}

//go:noinline
func Slider(g *tinygui.GuiContext, state *SliderState, props *SliderProps) bool {

	clicked := false
	colBackground := g.Theme.Background

	h := int16(30)

	eventAction, clickX, _ := HandleEvent(&g.Event, props.Disabled, true, false, props.X, props.Y, props.W, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & (Click | Drag | DragEnd)) != 0 {
		if eventAction == Drag {
			state.animateState = 11
		} else {
			state.animateState = 10
		}
		state.Value = float32(clickX-16) / float32(props.W-32)
		if state.Value < 0 {
			state.Value = 0
		} else if state.Value > 1 {
			state.Value = 1
		}
		clicked = true
		g.UpdateState()
	}

	tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, 16, h, colBackground)
	tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, props.W, 14, colBackground)
	tinydraw.FilledRectangleEx(g.Display, props.X, props.Y+19, props.W, 13, colBackground)
	tinydraw.FilledRectangleEx(g.Display, props.X+props.W-16, props.Y, 16, h, colBackground)

	//Slider bar

	sliderWidth := int16(state.Value * float32(props.W-32))
	sliderLeftCol := props.Color
	sliderRightCol := primitive.Desaturate(sliderLeftCol, 76)

	if state.animateState >= 0 {

		if state.animateState < 11 {
			g.InvalidateRect(props.X+sliderWidth+16-16, props.Y, props.X+sliderWidth+32, props.Y+h+1, false)
			state.animateState--
		}

		fillColor := g.Theme.Background
		if state.animateState > 0 {
			fillColor = primitive.Desaturate(sliderLeftCol, 51)
			fillColor = primitive.TransitionColor(fillColor, g.Theme.Background, uint8(state.animateState*23))
		}

		tinydraw.FilledCircleEx(g.Display, props.X+sliderWidth+16, props.Y+16, 15, fillColor)
	}

	primitive.FilledBox(g.Display, props.X+16, props.Y+14, sliderWidth+9, 5, sliderLeftCol)
	primitive.FilledBox(g.Display, props.X+sliderWidth+16, props.Y+14, props.W-sliderWidth-32, 5, sliderRightCol)

	tinydraw.FilledCircleEx(g.Display, props.X+sliderWidth+16, props.Y+16, 9, sliderLeftCol)

	return clicked
}

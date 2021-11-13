package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/primitive"

	"tinygo.org/x/tinydraw"
)

type ButtonStyle byte

const (
	Text ButtonStyle = iota
	Contained
	Outlined
)

type ButtonState struct {
	animateState int8
}

type ButtonProps struct {
	ComponentPos
	ComponentSize
	Label            string
	Style            ButtonStyle
	Disabled         bool
	Color            color.RGBA
	DisableElevation bool
	StartIcon        icon.Icon
	EndIcon          icon.Icon
}

// ATTEMPT 2

func NewButtonProps(g *tinygui.GuiContext) ButtonProps {
	return ButtonProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: -1,
			H: 30,
		},
		Label:            "Click",
		Style:            Text,
		Disabled:         false,
		DisableElevation: false,
		StartIcon:        nil,
		EndIcon:          nil,
		Color:            g.Theme.DefaultColor,
	}
}

//go:noinline
func btnText(g *tinygui.GuiContext, state *ButtonState, props *ButtonProps, w int16) color.RGBA {
	darkColor := props.Color
	//lightColor := g.Theme.Background
	if props.Disabled {
		darkColor = g.Theme.DefaultColor
		//lightColor = color.RGBA{166, 167, 168, 255}
	}

	//textColor := lightColor

	textColor := darkColor

	if state.animateState >= 0 {
		if state.animateState == 0 {
			darkColor = g.Theme.Background
		} else {
			darkColor = primitive.Desaturate(props.Color, 76)
			darkColor = primitive.TransitionColor(darkColor, g.Theme.Background, uint8(state.animateState*25))
		}

		primitive.FilledBox(g.Display, props.X, props.Y, w, props.H-2, darkColor)

		if state.animateState > 0 {
			g.InvalidateRect(props.X, props.Y, props.X+w, props.Y+props.H, false)
		}

		state.animateState--
	}

	return textColor

}

//go:noinline
func btnContained(g *tinygui.GuiContext, state *ButtonState, props *ButtonProps, w int16) color.RGBA {
	darkColor := props.Color
	lightColor := g.Theme.Background
	if props.Disabled {
		darkColor = g.Theme.DefaultColor
		lightColor = color.RGBA{166, 167, 168, 255}
	}

	fillColor := darkColor

	if state.animateState > 0 {
		fillColor = primitive.Desaturate(fillColor, 76)
		fillColor = primitive.TransitionColor(fillColor, darkColor, uint8(state.animateState*25))
		state.animateState--
		g.InvalidateRect(props.X, props.Y, props.X+w, props.Y+props.H, false)
	}

	shadow1 := g.Theme.Shadow1
	shadow2 := g.Theme.Shadow2
	if !props.DisableElevation && !props.Disabled {
		primitive.FilledBoxWithShadow(g.Display, props.X, props.Y, w, props.H, fillColor, shadow1, shadow2)
	} else {
		primitive.FilledBox(g.Display, props.X, props.Y, w, props.H-2, fillColor)
	}

	return lightColor
}

//go:noinline
func btnOutlined(g *tinygui.GuiContext, state *ButtonState, props *ButtonProps, w int16) color.RGBA {

	darkColor := props.Color
	if props.Disabled {
		darkColor = g.Theme.DefaultColor
	}

	if state.animateState > 0 {
		fillColor := darkColor
		if state.animateState == 1 {
			fillColor = g.Theme.Background
		} else {

			fillColor = primitive.Desaturate(fillColor, 76)
			fillColor = primitive.TransitionColor(fillColor, g.Theme.Background, uint8(state.animateState*25))
		}

		tinydraw.FilledRectangleEx(g.Display, int16(props.X+1), int16(props.Y+1), int16(w-1), int16(props.H-3), fillColor)

		g.InvalidateRect(props.X, props.Y, props.X+w, props.Y+props.H, false)

		state.animateState--
	}

	//-2 for the shadow on contained buttons
	primitive.OutlineBox(g.Display, props.X, props.Y, w, props.H-2, darkColor)

	return darkColor

}

//go:noinline
func btnLabel(g *tinygui.GuiContext, state *ButtonState, props *ButtonProps, textAndIconWidth int16, textColor color.RGBA, centreY int16, w int16) int16 {

	if props.StartIcon != nil {
		textAndIconWidth += 20
	}

	if props.EndIcon != nil {
		textAndIconWidth += 20
	}

	textStartXOffset := (w / 2) - (textAndIconWidth / 2)

	if props.StartIcon != nil {
		textStartXOffset += 20
	}

	if props.Label != "" {
		primitive.WriteLine(g.Display, g.Theme.Font, props.X+textStartXOffset, props.Y+(centreY+5), props.Label, textColor)
	}

	return textStartXOffset
}

//go:noinline
func btnIcons(g *tinygui.GuiContext, state *ButtonState, props *ButtonProps, textStartXOffset int16, ow int16, textColor color.RGBA, centreY int16) {

	if props.StartIcon != nil {
		props.StartIcon(g, props.X+textStartXOffset-17, props.Y+(centreY-9), 15, textColor)
	}

	if props.EndIcon != nil {
		props.EndIcon(g, props.X+textStartXOffset+int16(ow)+2, props.Y+(centreY-9), 15, textColor)
	}

}

//go:noinline
func Button(g *tinygui.GuiContext, state *ButtonState, props *ButtonProps) bool {

	ow := uint32(0)
	if props.Label != "" {
		_, ow = primitive.LineWidth(g.Theme.Font, props.Label)
	}

	w := props.W

	if w == -1 {
		w = int16(ow) + 10
		if props.StartIcon != nil {
			w += 20
		}
		if props.EndIcon != nil {
			w += 20
		}
	}

	clicked := false

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, true, props.X, props.Y, w, props.H)

	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.animateState = 10
		clicked = true
		g.UpdateState()
	}
	textColor := g.Theme.Text

	switch props.Style {
	case Text:
		textColor = btnText(g, state, props, w)
	case Contained:
		textColor = btnContained(g, state, props, w)

	case Outlined:
		textColor = btnOutlined(g, state, props, w)
	}

	centreY := (props.H / 2)
	textStartXOffset := btnLabel(g, state, props, int16(ow), textColor, centreY, w)
	btnIcons(g, state, props, textStartXOffset, int16(ow), textColor, centreY)

	return clicked

}

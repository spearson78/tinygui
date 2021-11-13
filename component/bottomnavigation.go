package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/primitive"
)

type BottomNavigationState struct {
	animateState int8
	lastSelected uint8
	Selected     uint8
}

type BottomNavigationOption struct {
	Icon  icon.Icon
	Label string
}

type BottomNavigationProps struct {
	ComponentPos
	ComponentSize
	Options    []BottomNavigationOption
	Disabled   bool
	Color      color.RGBA
	PermaLabel bool
}

var defaultBottomNavOptions = []BottomNavigationOption{
	{Icon: icon.Calc, Label: "Calc"},
	{Icon: icon.Timeline, Label: "Data"},
	{Icon: icon.Announcement, Label: "Test"},
}

func NewBottomNavigationProps(g *tinygui.GuiContext) BottomNavigationProps {
	displayWidth, _ := g.Display.Size()

	return BottomNavigationProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: displayWidth,
			H: -1,
		},
		Options:    defaultBottomNavOptions,
		Disabled:   false,
		Color:      g.Theme.DefaultColor,
		PermaLabel: false,
	}
}

func BottomNavigation(g *tinygui.GuiContext, state *BottomNavigationState, props *BottomNavigationProps) bool {

	optionWidth := props.W / int16(len(props.Options))

	clicked := false

	h := int16(40)

	draw := true

	eventAction, clickX, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, props.W, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.animateState = 10
		state.Selected = uint8(clickX / optionWidth)
		clicked = true
		g.UpdateState()
	}
	if (eventAction & Update) != 0 {
		if state.lastSelected == state.Selected {
			draw = false
		}
	}

	state.lastSelected = state.Selected

	if draw {

		x := props.X
		iconWidth := int16(16)
		iconOffset := (optionWidth / 2) - (iconWidth / 2)

		for i, option := range props.Options {

			col := g.Theme.Border

			if i == int(state.Selected) {

				col = props.Color

				if state.animateState > 0 {
					fillColor := primitive.Desaturate(props.Color, 76)
					fillColor = primitive.TransitionColor(fillColor, g.Theme.Background, uint8(state.animateState*25))

					g.Display.FillRect(x, props.Y, optionWidth, h, fillColor)

					if state.animateState > 0 {
						g.InvalidateRect(x, props.Y, x+optionWidth, props.Y+h, false)
					}

					state.animateState--
				} else {
					g.Display.FillRect(x, props.Y, optionWidth, h, g.Theme.Background)
				}
			} else {
				g.Display.FillRect(x, props.Y, optionWidth, h, g.Theme.Background)
			}

			if props.PermaLabel || i == int(state.Selected) {
				option.Icon(g, x+iconOffset, props.Y+5, 16, col)
				_, ow := primitive.LineWidth(g.Theme.Font, option.Label)
				textOffset := (optionWidth / 2) - int16(ow/2)
				primitive.WriteLine(g.Display, g.Theme.Font, x+textOffset, props.Y+35, option.Label, col)
			} else {
				option.Icon(g, x+iconOffset, props.Y+15, 16, col)
			}

			x = x + optionWidth
		}
	}

	return clicked

}

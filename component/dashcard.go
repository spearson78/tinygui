package component

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type DashCardState struct {
	animateState int8
	Value        string
	ActionIcon   icon.Icon
	ActionLabel  string
}

type DashCardProps struct {
	ComponentPos
	ComponentSize
	BadgeIcon icon.Icon
	Label     string
	Disabled  bool
	Color     color.RGBA
}

func NewDashCardProps(g *tinygui.GuiContext) DashCardProps {
	return DashCardProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: 110,
			H: -1,
		},
		Label:     "Click",
		Disabled:  false,
		BadgeIcon: icon.Announcement,
		Color:     g.Theme.DefaultColor,
	}
}

//go:noinline
func DashCard(g *tinygui.GuiContext, state *DashCardState, props *DashCardProps) bool {

	_, labelWidth := primitive.LineWidth(g.Theme.SecondaryFont, props.Label)
	_, valueWidth := primitive.LineWidth(g.Theme.Font, state.Value)

	clicked := false

	h := int16(70)

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, props.W, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		state.animateState = 10
		clicked = true
		g.UpdateState()
	}

	border := g.Theme.Border
	white := g.Theme.Background

	if (g.Event.Type & event.Invalidate) != 0 {
		tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, props.W, h, white)

		primitive.OutlineBoxWithShadow(g.Display, props.X, props.Y+10, props.W, h-10, border, g.Theme.Shadow1, g.Theme.Shadow2)

		primitive.FilledBoxWithShadow(g.Display, props.X+10, props.Y, 40, 42, props.Color, g.Theme.Shadow1, g.Theme.Shadow2)

		tinydraw.LineEx(g.Display, props.X+10, props.Y+48, props.X+props.W-10, props.Y+48, border)
		props.BadgeIcon(g, props.X+20, props.Y+10, 20, white)
		primitive.WriteLine(g.Display, g.Theme.SecondaryFont, props.X+props.W-int16(labelWidth)-10, props.Y+18, props.Label, border)

	} else {
		tinydraw.FilledRectangleEx(g.Display, props.X+10+41, props.Y+24, props.W-52, 17, g.Theme.Background)
	}

	buttonBack := g.Theme.Background
	if state.animateState > 0 {

		buttonBack = primitive.Desaturate(border, 76)
		buttonBack = primitive.TransitionColor(buttonBack, g.Theme.Background, uint8(state.animateState*25))
	}

	tinydraw.FilledRectangleEx(g.Display, props.X+1, props.Y+49, props.W-2, 18, buttonBack)

	if state.animateState > 0 {
		g.InvalidateRect(props.X+1, props.Y+49, props.X+1+props.W-2, props.Y+49+18, false)
		state.animateState--
	}

	primitive.WriteLine(g.Display, g.Theme.Font, props.X+props.W-int16(valueWidth)-10, props.Y+38, state.Value, g.Theme.Text)
	state.ActionIcon(g, props.X+11, props.Y+52, 10, border)
	primitive.WriteLine(g.Display, g.Theme.SecondaryFont, props.X+11+10+3, props.Y+60, state.ActionLabel, border)

	return clicked
}

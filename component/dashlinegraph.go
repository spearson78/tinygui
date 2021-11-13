package component

import (
	"image/color"
	"strconv"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type DashLineGraphState struct {
	Values        []float32
	SubTitle      string
	SubTitleColor color.RGBA
	//ActionIcon  IconFunc
	//ActionLabel string
}

type DashLineGraphProps struct {
	ComponentPos
	ComponentSize
	Label    string
	XLabels  []string
	YStart   int16
	YEnd     int16
	YStep    int16
	Disabled bool
	Color    color.RGBA
}

var defaultDashLineGraphXLabels = []string{"0", "10", "20", "30", "40", "50", "60", "70", "80", "90", "100"}

func NewDashLineGraphProps(g *tinygui.GuiContext) DashLineGraphProps {
	return DashLineGraphProps{
		ComponentPos: ComponentPos{
			X: 10,
			Y: 10,
		},
		ComponentSize: ComponentSize{
			W: 220,
			H: -1,
		},
		Disabled: false,
		Color:    g.Theme.DefaultColor,
		YStart:   0,
		YEnd:     100,
		YStep:    10,
		Label:    "Line Graph",
		XLabels:  defaultDashLineGraphXLabels,
	}
}

//go:noinline
func DashLineGraph(g *tinygui.GuiContext, state *DashLineGraphState, props *DashLineGraphProps) bool {

	border := g.Theme.Border
	white := g.Theme.Background
	axisColor := white
	lineColor := white

	graphStartY := int16(118)
	graphHeight := int16(110)
	graphYScale := float32(graphHeight) / (float32(props.YEnd) - float32(props.YStart))
	h := graphStartY + 40

	eventAction, _, _ := HandleEvent(&g.Event, props.Disabled, false, false, props.X, props.Y, props.W, h)
	if eventAction == Ignore {
		return false
	}
	if (eventAction & Click) != 0 {
		return false
	}

	graphStartX := int16(26)
	graphWidth := props.W - graphStartX - 15
	grapXScale := float32(graphWidth) / float32(len(props.XLabels)-1)

	if (g.Event.Type & event.Invalidate) != 0 {
		tinydraw.FilledRectangleEx(g.Display, props.X, props.Y, props.W, h, white)

		primitive.OutlineBoxWithShadow(g.Display, props.X, props.Y+10, props.W, h, border, g.Theme.Shadow1, g.Theme.Shadow2)

		primitive.FilledBoxWithShadow(g.Display, props.X+10, props.Y, props.W-20, 130, props.Color, g.Theme.Shadow1, g.Theme.Shadow2)

		yCounter := float32(0)
		for yLabelInt := props.YStart; yLabelInt <= props.YEnd; yLabelInt += props.YStep {
			yLabelOffset := yCounter * graphYScale
			yLabelStr := strconv.FormatInt(int64(yLabelInt), 10)

			primitive.WriteLine(g.Display, g.Theme.SecondaryFont, props.X+12, props.Y+graphStartY-int16(yLabelOffset), yLabelStr, axisColor)
			yCounter += float32(props.YStep)
		}

		for x := 0; x < len(props.XLabels); x++ {
			xLabelOffset := int16(float32(x) * grapXScale)
			xLabelStr := props.XLabels[x]

			primitive.WriteLine(g.Display, g.Theme.SecondaryFont, props.X+graphStartX+xLabelOffset, props.Y+graphStartY+7, xLabelStr, axisColor)
		}

		primitive.WriteLine(g.Display, g.Theme.Font, props.X+10, props.Y+graphStartY+30, props.Label, g.Theme.Text)

	} else {
		tinydraw.FilledRectangleEx(g.Display, props.X+graphStartX-2, props.Y+graphStartY-graphHeight-2, graphWidth+5, graphHeight+5, props.Color)
		tinydraw.FilledRectangleEx(g.Display, props.X+10, props.Y+graphStartY+34, props.W-20, 7, g.Theme.Background)

	}

	tinydraw.LineEx(g.Display, props.X+graphStartX, props.Y+graphStartY, props.X+graphStartX, props.Y+graphStartY-graphHeight, axisColor)
	tinydraw.LineEx(g.Display, props.X+graphStartX, props.Y+graphStartY, props.X+graphStartX+graphWidth, props.Y+graphStartY, axisColor)

	lastXPos := int16(0)
	lastYPos := int16(0)
	for x := 0; x < len(state.Values); x++ {
		xPos := props.X + graphStartX + int16(float32(x)*grapXScale)
		yPos := props.Y + graphStartY - int16(float32(state.Values[x]-float32(props.YStart))*graphYScale)
		tinydraw.FilledCircleEx(g.Display, xPos, yPos, 2, lineColor)

		if x > 0 {
			tinydraw.LineEx(g.Display, lastXPos, lastYPos, xPos, yPos, lineColor)
		}

		lastXPos = xPos
		lastYPos = yPos
	}

	primitive.WriteLine(g.Display, g.Theme.SecondaryFont, props.X+10, props.Y+graphStartY+40, state.SubTitle, state.SubTitleColor)

	return false
}

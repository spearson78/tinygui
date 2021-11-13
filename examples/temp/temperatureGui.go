package temp

import (
	"math"
	"strconv"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/component"
	"github.com/spearson78/tinygui/icon"
	"github.com/spearson78/tinygui/theme"
)

var txtTempDisplayState = component.DashCardState{
	Value:       "0",
	ActionIcon:  icon.Edit,
	ActionLabel: "Fahrenheit",
}

var lineGraphState = component.DashLineGraphState{
	Values: []float32{24, 24, 24, 24, 24, 24, 24},
}

var tempCelsius = true
var lastTemp = float32(20)
var graphUpdateCounter = int64(0)

//go:noinline
func DoSensor(g *tinygui.TinyGui, t *theme.Theme, readTemp func() float32) {
	temp := readTemp()

	if math.Abs(float64(lastTemp)-float64(temp)) >= 1.0 {

		t := float64(temp)
		suffix := "°C"
		if !tempCelsius {
			t = (t * (9 / 5)) + 32
			suffix = "°F"
		}

		tempStr := strconv.FormatFloat(t, 'f', 0, 64)
		txtTempDisplayState.Value = tempStr + suffix
		g.UpdateState()
		lastTemp = temp

	}

	graphUpdateCounter++
	if graphUpdateCounter%10 == 0 {

		lineGraphState.Values[0] = lineGraphState.Values[1]
		lineGraphState.Values[1] = lineGraphState.Values[2]
		lineGraphState.Values[2] = lineGraphState.Values[3]
		lineGraphState.Values[3] = lineGraphState.Values[4]
		lineGraphState.Values[4] = lineGraphState.Values[5]
		lineGraphState.Values[5] = lineGraphState.Values[6]
		lineGraphState.Values[6] = temp

		percentChange := ((lineGraphState.Values[6] - lineGraphState.Values[0]) / lineGraphState.Values[6]) * float32(100.0)
		if percentChange >= 0.0 {
			lineGraphState.SubTitle = strconv.FormatFloat(float64(percentChange), 'f', 2, 64) + "% Increase"
			lineGraphState.SubTitleColor = t.SuccessColor
		} else {
			lineGraphState.SubTitle = strconv.FormatFloat(math.Abs(float64(percentChange)), 'f', 2, 64) + "% Decrease"
			lineGraphState.SubTitleColor = t.ErrorColor
		}

		g.UpdateState()

	}
}

var tempXLabels = []string{"6", "5", "4", "3", "2", "1", "0"}

//go:noinline
func TemperatureSensorGui(g *tinygui.GuiContext) {

	dashCardProps := component.NewDashCardProps(g)
	dashCardProps.X = 5
	dashCardProps.Y = 10
	dashCardProps.Label = "Temp"

	if component.DashCard(g, &txtTempDisplayState, &dashCardProps) {
		tempCelsius = !tempCelsius
		if tempCelsius {
			txtTempDisplayState.ActionLabel = "Fahrenheit"
		} else {
			txtTempDisplayState.ActionLabel = "Celsius"
		}
		lastTemp = 0
	}

	dashLineProps := component.NewDashLineGraphProps(g)
	dashLineProps.X = 10
	dashLineProps.Y = 85
	dashLineProps.Color = g.Theme.SuccessColor
	dashLineProps.YStart = 16
	dashLineProps.YEnd = 30
	dashLineProps.YStep = 2
	dashLineProps.Label = "Temperature"
	dashLineProps.XLabels = tempXLabels
	if component.DashLineGraph(g, &lineGraphState, &dashLineProps) {
	}

}

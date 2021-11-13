package nav

import (
	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/component"
	"github.com/spearson78/tinygui/examples/calc"
	"github.com/spearson78/tinygui/examples/temp"
	"github.com/spearson78/tinygui/examples/test"
)

var NavigationState = component.BottomNavigationState{
	Selected: 0,
}

func NavGui(g *tinygui.GuiContext) {

	switch NavigationState.Selected {
	case 0:
		calc.CalcGui(g)
	case 1:
		temp.TemperatureSensorGui(g)
	default:
		test.TestGui(g)
	}

	navProps := component.NewBottomNavigationProps(g)
	navProps.X = 0
	navProps.Y = 280
	navProps.PermaLabel = false
	if component.BottomNavigation(g, &NavigationState, &navProps) {
		g.Invalidate(true)
	}
}

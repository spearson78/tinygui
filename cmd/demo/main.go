package main

import (
	"time"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/examples/calibrate"
	"github.com/spearson78/tinygui/examples/nav"
	"github.com/spearson78/tinygui/examples/temp"
	"github.com/spearson78/tinygui/theme"
)

func guiFunc(g *tinygui.GuiContext) {
	if !calibrate.IsCalibrated() {
		calibrate.Gui(g)
		if calibrate.IsCalibrated() {
			g.TriggerInvalidate = true
			g.ClearDisplay = true
		}
	} else {
		nav.NavGui(g)
		//toggleled.ToggleLedGui(g)
	}
}

func main() {

	//stack.UpdateStackPointer("main")
	//baseStackPointer := stack.MinStackPointer

	display, touch := initHardware()

	theme := theme.DefaultTheme()
	gui := tinygui.New(display, touch, theme, calibrate.TouchCalibration)

	gui.Init()

	for {
		time.Sleep(50 * time.Millisecond)
		g := gui.DoTouchNonBlocking()
		if g != nil {
			guiFunc(g)
		}

		if nav.NavigationState.Selected == 1 {
			temp.DoSensor(gui, &theme, ReadTemp)
		}

		/*
			txt := strconv.FormatInt(int64(baseStackPointer)-int64(stack.MinStackPointer), 10)
			primitive.WriteLine(display, theme.Font, 10, 40, txt, color.RGBA{255, 0, 0, 0255})
			primitive.WriteLine(display, theme.Font, 10, 70, stack.MinStackLocation, color.RGBA{255, 0, 0, 0255})
		*/
	}

}

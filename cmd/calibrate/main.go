package main

import (
	"time"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/examples/calibrate"
	"github.com/spearson78/tinygui/theme"
)

func main() {

	display, touch := initHardware()

	gui := tinygui.New(display, touch, theme.DefaultTheme(), calibrate.TouchCalibration)

	gui.Init()

	for {
		time.Sleep(50 * time.Millisecond)
		g := gui.DoTouchNonBlocking()
		if g != nil {
			calibrate.Gui(g)
		}
	}

}

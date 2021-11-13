package calibrate

import (
	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/primitive"
	"tinygo.org/x/tinydraw"
)

type calibrationState struct {
	state        uint8
	topLeftX     uint16
	topLeftY     uint16
	bottomRightX uint16
	bottomRightY uint16

	xPixelSize uint16
	xOffset    uint16
	yPixelSize uint16
	yOffset    uint16
}

func IsCalibrated() bool {
	return CalibrationState.state == 2
}

func TouchCalibration(x int, y int) (int32, int32) {
	if CalibrationState.state != 2 {
		return int32(x), int32(y)
	} else {
		rx := (int32(x) / int32(CalibrationState.xPixelSize)) - int32(CalibrationState.xOffset)
		ry := (int32(y) / int32(CalibrationState.yPixelSize)) - int32(CalibrationState.yOffset)

		return rx, ry
	}
}

var CalibrationState = calibrationState{}

func handleCalibrateEvent(g *tinygui.GuiContext, evt event.Event, px, py int16) (rx, ry uint16, clicked bool) {
	fillColor := g.Theme.DefaultColor

	if (evt.Type & (event.DragEnd | event.Click)) != 0 {
		fillColor = g.Theme.SuccessColor
		clicked = true
		rx = evt.X
		ry = evt.Y
	}

	tinydraw.FilledCircleEx(g.Display, px, py, 10, fillColor)

	return
}

//go:noinline
func Gui(g *tinygui.GuiContext) {

	w, h := g.Display.Size()

	switch CalibrationState.state {
	case 0:
		px := int16(20)
		py := int16(20)
		rx, ry, clicked := handleCalibrateEvent(g, g.Event, px, py)
		if clicked {
			CalibrationState.state++
			CalibrationState.topLeftX = rx
			CalibrationState.topLeftY = ry
			g.UpdateState()
		}
	case 1:
		px := w - 20
		py := h - 20
		rx, ry, clicked := handleCalibrateEvent(g, g.Event, px, py)
		if clicked {
			CalibrationState.state++
			CalibrationState.bottomRightX = rx
			CalibrationState.bottomRightY = ry

			CalibrationState.xPixelSize = uint16((int32(CalibrationState.bottomRightX) - int32(CalibrationState.topLeftX)) / int32(w-40))
			CalibrationState.xOffset = ((CalibrationState.topLeftX) / CalibrationState.xPixelSize) - 20

			CalibrationState.yPixelSize = uint16((int32(CalibrationState.bottomRightY) - int32(CalibrationState.topLeftY)) / int32(h-40))
			CalibrationState.yOffset = ((CalibrationState.topLeftY) / CalibrationState.yPixelSize) - 20

		}
	case 2:
		if (g.Event.Type & event.Click) != 0 {
			tinydraw.FilledCircleEx(g.Display, int16(g.Event.X), int16(g.Event.Y), 10, g.Theme.SecondaryColor)
		}

	}

	if (g.Event.Type & event.Invalidate) != 0 {
		txt := "Touch Calibration"
		_, ow := primitive.LineWidth(g.Theme.Font, txt)

		txtX := (w / 2) - (int16(ow) / 2)
		txtY := (h / 2) + (int16(g.Theme.Font.Ysize) / 2)

		primitive.WriteLine(g.Display, g.Theme.Font, txtX, txtY, txt, g.Theme.Text)
	}

}

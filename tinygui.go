package tinygui

import (
	"image"

	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/theme"

	"tinygo.org/x/drivers/touch"
)

type Component func(g *GuiContext)

type TouchCalibration func(x int, y int) (int32, int32)

type TouchSource interface {
	Read() (int32, int32, bool)
}

type TinyGui struct {
	guiContext       GuiContext
	clip             ClippedDisplay
	tch              touch.Pointer
	touchCalibration TouchCalibration

	touchInProgress bool
	dragInProgress  bool
	touchStartX     int32
	touchStartY     int32
	touchLastX      int32
	touchLastY      int32
}

func New(d Displayer, t touch.Pointer, theme theme.Theme, touchCalibration TouchCalibration) *TinyGui {
	tinyGui := TinyGui{
		guiContext: GuiContext{
			Theme:             theme,
			TriggerInvalidate: true,
			ClearDisplay:      true,
			InvalidX0:         -1,
			InvalidY0:         -1,
			InvalidX1:         -1,
			InvalidY1:         -1,
		},
		clip:             ClippedDisplay{D: d},
		tch:              t,
		touchCalibration: touchCalibration,
	}

	tinyGui.guiContext.Display = &tinyGui.clip

	return &tinyGui
}

func (t *TinyGui) Init() {
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

const dragSensitivity = int32(3)

func (t *TinyGui) readTouch() (int32, int32, bool) {
	tp := t.tch.ReadTouchPoint()
	touching := tp.Z != 0

	if touching {
		x, y := t.touchCalibration(tp.X, tp.Y)
		return x, y, touching
	}
	return 0, 0, false
}

func (t *TinyGui) UpdateState() {
	t.guiContext.TriggerUpdateState = true
}

//go:noinline
func (t *TinyGui) DoTouchNonBlocking() *GuiContext {
	t.clip.ClipRect = image.Rectangle{}

	t.guiContext.Event.Type = event.None

	if t.touchInProgress {

		//Continue Drag
		x, y, touching := t.readTouch()

		if !touching {
			if t.dragInProgress {
				t.guiContext.Event.Type |= event.DragEnd
			} else {
				t.guiContext.Event.Type |= event.Click
			}
			t.touchInProgress = false
			t.dragInProgress = false
		} else {

			deltaX := abs(t.touchLastX - x)
			deltaY := abs(t.touchLastY - y)
			dragSens := dragSensitivity
			if !t.dragInProgress {
				dragSens *= 3
			}
			if deltaX > dragSens || deltaY > dragSens {
				t.dragInProgress = true
				t.guiContext.Event.Type |= event.Drag
				t.touchLastX = x
				t.touchLastY = y
			}
		}

		if t.guiContext.Event.Type != event.None {

			t.guiContext.Event.X = uint16(t.touchStartX)
			t.guiContext.Event.Y = uint16(t.touchStartY)
			t.guiContext.Event.DragX = uint16(t.touchLastX)
			t.guiContext.Event.DragY = uint16(t.touchLastY)
		}
	} else {

		x, y, touching := t.readTouch()
		if touching {
			//Start of touch maybe a click
			t.touchInProgress = true
			t.touchStartX = x
			t.touchStartY = y
			t.touchLastX = x
			t.touchLastY = y
		}
	}

	if t.guiContext.TriggerUpdateState {
		t.guiContext.TriggerUpdateState = false
		if t.guiContext.TriggerInvalidate && t.guiContext.InvalidX0 == -1 {
			t.guiContext.TriggerInvalidate = false
			t.guiContext.Event.Type |= event.Invalidate

			w, h := t.clip.D.Size()
			if t.guiContext.ClearDisplay {
				t.clip.FillRect(0, 0, w-1, h-1, t.guiContext.Theme.Background)
				t.guiContext.ClearDisplay = false
			}

		} else {
			t.guiContext.Event.Type |= event.Update
		}

	} else if t.guiContext.TriggerInvalidate && (t.guiContext.InvalidX0 == -1 || t.guiContext.Event.Type == event.None) {

		w, h := t.clip.D.Size()
		if t.guiContext.InvalidX0 != -1 {
			t.clip.ClipRect = image.Rect(int(t.guiContext.InvalidX0), int(t.guiContext.InvalidY0), int(t.guiContext.InvalidX1+1), int(t.guiContext.InvalidY1+1))
			t.guiContext.InvalidX0 = -1
			t.guiContext.InvalidY0 = -1
			t.guiContext.InvalidX1 = -1
			t.guiContext.InvalidY1 = -1
		}

		if t.guiContext.ClearDisplay {
			t.clip.FillRect(0, 0, w-1, h-1, t.guiContext.Theme.Background)
			t.guiContext.ClearDisplay = false
		}

		t.guiContext.TriggerInvalidate = false
		t.guiContext.Event.Type |= event.Invalidate
	}

	if t.guiContext.Event.Type != event.None {
		return &t.guiContext
	} else {
		return nil
	}
}

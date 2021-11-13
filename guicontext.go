package tinygui

import (
	"math"

	"github.com/spearson78/tinygui/event"
	"github.com/spearson78/tinygui/theme"
)

type GuiContext struct {
	Display Displayer
	Theme   theme.Theme
	Event   event.Event

	TriggerUpdateState bool
	TriggerInvalidate  bool
	ClearDisplay       bool

	InvalidX0 int16
	InvalidY0 int16
	InvalidX1 int16
	InvalidY1 int16
}

func (t *GuiContext) UpdateState() {
	t.TriggerUpdateState = true
}

func (t *GuiContext) Invalidate(clearDisplay bool) {
	t.TriggerInvalidate = true
	t.InvalidX0 = 0
	t.InvalidY0 = 0
	t.InvalidX1 = math.MaxInt16 - 1
	t.InvalidY1 = math.MaxInt16 - 1
	t.ClearDisplay = t.ClearDisplay || clearDisplay
}

func minInvalidCoord(cur int16, new int16) int16 {
	if cur == -1 || new < cur {
		return new
	} else {
		return cur
	}
}

func maxInvalidCoord(cur int16, new int16) int16 {
	if cur == -1 || new > cur {
		return new
	} else {
		return cur
	}
}

func (t *GuiContext) InvalidateRect(x0, y0, x1, y1 int16, clearDisplay bool) {
	t.TriggerInvalidate = true
	t.ClearDisplay = t.ClearDisplay || clearDisplay
	t.InvalidX0 = minInvalidCoord(t.InvalidX0, x0)
	t.InvalidY0 = minInvalidCoord(t.InvalidY0, y0)
	t.InvalidX1 = maxInvalidCoord(t.InvalidX1, x1)
	t.InvalidY1 = maxInvalidCoord(t.InvalidY1, y1)
}

package component

import (
	"github.com/spearson78/tinygui/event"
)

type EventAction uint8

const (
	Ignore     EventAction = 0
	Invalidate EventAction = 1
	Update     EventAction = 2
	Click      EventAction = 4
	Drag       EventAction = 8
	DragEnd    EventAction = 16
)

//go:noinline
func HandleEvent(guiEvent *event.Event, disabled bool, drag bool, noExternalState bool, x, y, w, h int16) (eventAction EventAction, rx int16, ry int16) {

	//stack.UpdateStackPointer("HandleEvent")

	if (guiEvent.Type & event.Click) != 0 {
		if !disabled {
			if int16(guiEvent.X) > x && int16(guiEvent.X) < (x+w) && int16(guiEvent.Y) > y && int16(guiEvent.Y) < (y+h) {
				eventAction |= Click
				rx = int16(guiEvent.X) - x
				ry = int16(guiEvent.Y) - y
			}
		}
	} else if (guiEvent.Type & (event.Drag | event.DragEnd)) != 0 {

		if drag {
			if int16(guiEvent.X) > x && int16(guiEvent.X) < (x+w) && int16(guiEvent.Y) > y && int16(guiEvent.Y) < (y+h) {

				dragType := Drag
				if (guiEvent.Type & event.DragEnd) != 0 {
					dragType = DragEnd
				}

				eventAction |= dragType
				rx = int16(guiEvent.DragX) - x
				ry = int16(guiEvent.DragY) - y
			}
		}
	}

	if (guiEvent.Type & event.Invalidate) != 0 {
		eventAction |= Invalidate
	}

	if (guiEvent.Type & event.Update) != 0 {
		if !noExternalState {
			eventAction |= Update
		}
	}

	return
}

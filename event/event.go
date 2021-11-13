package event

type Event struct {
	Type Type

	X uint16
	Y uint16

	DragX uint16
	DragY uint16
}

package event

type Type byte

const (
	None       Type = 0
	Invalidate Type = 1
	Update     Type = 2
	Click      Type = 4
	Drag       Type = 8
	DragEnd    Type = 16
)

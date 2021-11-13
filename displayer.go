package tinygui

import (
	"image/color"

	"tinygo.org/x/drivers"
)

type Displayer interface {
	drivers.Displayer
	FillRect(x, y, w, h int16, c color.RGBA)
}

package adapter

import (
	"image/color"

	"tinygo.org/x/drivers"
)

type fillRectAdapter struct {
	d drivers.Displayer
}

func AddFillRect(d drivers.Displayer) fillRectAdapter {
	return fillRectAdapter{
		d: d,
	}
}

func (d *fillRectAdapter) Size() (x, y int16) {
	return d.d.Size()
}

func (d *fillRectAdapter) SetPixel(x, y int16, c color.RGBA) {
	d.d.SetPixel(x, y, c)
}

func (d *fillRectAdapter) Display() error {
	return d.d.Display()
}

func (d *fillRectAdapter) FillRect(x, y, w, h int16, c color.RGBA) {
	for py := y; py < y+h; py++ {
		for px := x; px < x+w; px++ {
			d.d.SetPixel(px, py, c)
		}
	}
}

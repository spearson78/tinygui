//go:build !tinygo
// +build !tinygo

package adapter

import (
	"image/color"

	"github.com/sago35/tinydisplay"
)

type tinyDisplayAdapter struct {
	c *tinydisplay.Client
}

func NewTinyDisplayAdapter(c *tinydisplay.Client) tinyDisplayAdapter {
	return tinyDisplayAdapter{
		c: c,
	}
}

func (d *tinyDisplayAdapter) Size() (x, y int16) {
	return d.c.Size()
}

func (d *tinyDisplayAdapter) SetPixel(x, y int16, c color.RGBA) {
	d.c.SetPixel(x, y, c)
}

func (d *tinyDisplayAdapter) Display() error {
	return d.c.Display()
}

func (d *tinyDisplayAdapter) FillRect(x, y, w, h int16, c color.RGBA) {
	d.c.FillRectangle(x, y, w, h, c)
}

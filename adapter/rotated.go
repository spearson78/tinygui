package adapter

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"tinygo.org/x/drivers/touch"
)

type rotateAdapter struct {
	d  tinygui.Displayer
	t  touch.Pointer
	xs int16
	ys int16
}

func Rotate(d tinygui.Displayer, t touch.Pointer) rotateAdapter {
	xs, ys := d.Size()
	return rotateAdapter{
		d:  d,
		t:  t,
		xs: xs,
		ys: ys,
	}
}

func (r *rotateAdapter) Size() (int16, int16) {
	return r.ys, r.xs
}

func (r *rotateAdapter) SetPixel(x, y int16, c color.RGBA) {
	r.d.SetPixel(r.xs-y, x, c)
}

func (r *rotateAdapter) Display() error {
	return r.d.Display()
}

func (r *rotateAdapter) FillRect(x, y, w, h int16, c color.RGBA) {
	r.d.FillRect(r.xs-y-h, x, h, w, c)
}

func (r *rotateAdapter) Read() touch.Point {
	p := r.t.ReadTouchPoint()
	return touch.Point{
		X: p.Y,
		Y: int(r.xs) - p.X,
		Z: p.Z,
	}
}

package adapter

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"tinygo.org/x/drivers/touch"
)

type upsideDownAdapter struct {
	d  tinygui.Displayer
	t  touch.Pointer
	xs int16
	ys int16
}

func UpsideDown(d tinygui.Displayer, t touch.Pointer) upsideDownAdapter {
	xs, ys := d.Size()
	return upsideDownAdapter{
		d:  d,
		t:  t,
		xs: xs,
		ys: ys,
	}
}

func (r *upsideDownAdapter) Size() (int16, int16) {
	return r.xs, r.ys
}

func (r *upsideDownAdapter) SetPixel(x, y int16, c color.RGBA) {
	r.d.SetPixel(r.xs-x, r.ys-y, c)
}

func (r *upsideDownAdapter) Display() error {
	return r.d.Display()
}

func (r *upsideDownAdapter) FillRect(x, y, w, h int16, c color.RGBA) {
	r.d.FillRect(r.xs-x-w, r.ys-y-h, w, h, c)
}

func (r *upsideDownAdapter) ReadTouchPoint() touch.Point {
	tp := r.t.ReadTouchPoint()
	tp.X = ((1 << 16) - tp.X)
	tp.Y = ((1 << 16) - tp.Y)
	return tp
}

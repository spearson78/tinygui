package tinygui

import (
	"image"
	"image/color"
)

type ClippedDisplay struct {
	D        Displayer
	ClipRect image.Rectangle
}

func (r *ClippedDisplay) Size() (int16, int16) {
	return r.D.Size()
}

func (r *ClippedDisplay) SetPixel(x, y int16, c color.RGBA) {

	if r.ClipRect.Empty() {
		r.D.SetPixel(x, y, c)
	} else {
		if x < int16(r.ClipRect.Min.X) || x > int16(r.ClipRect.Max.X) || y < int16(r.ClipRect.Min.Y) || y > int16(r.ClipRect.Max.Y) {
			return
		}

		r.D.SetPixel(x, y, c)
	}
}

func (r *ClippedDisplay) Display() error {
	return r.D.Display()
}

func (r *ClippedDisplay) FillRect(x, y, w, h int16, c color.RGBA) {

	if r.ClipRect.Empty() {
		r.D.FillRect(x, y, w, h, c)
	} else {
		intersection := r.ClipRect.Intersect(image.Rect(int(x), int(y), int(x+w), int(y+h)))

		if !intersection.Empty() {
			r.D.FillRect(int16(intersection.Min.X), int16(intersection.Min.Y), int16(intersection.Dx()), int16(intersection.Dy()), c)
		}
	}
}

package primitive

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"tinygo.org/x/tinydraw"
)

type Point8 struct {
	X uint8
	Y uint8
}

type LineSegment struct {
	A Point8
	B Point8
}

type Shape struct {
	Lines []LineSegment
}

func getLineIntersection(x0, y0, x1, y1, x2, y2, x3, y3 float32) (float32, bool) {

	s1_x := x1 - x0
	s1_y := y1 - y0
	s2_x := x3 - x2
	s2_y := y3 - y2

	s := (-s1_y*(x0-x2) + s1_x*(y0-y2)) / (-s2_x*s1_y + s1_x*s2_y)
	t := (s2_x*(y0-y2) - s2_y*(x0-x2)) / (-s2_x*s1_y + s1_x*s2_y)

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		// Collision detected
		x := x0 + (t * s1_x)
		return x, true
	} else {
		return 0, false
	}
}

//go:noinline
func FillShape(d tinygui.Displayer, x, y, w, h int16, s *Shape, c color.RGBA) {

	wscale := float32(255) / float32(w)
	hscale := float32(255) / float32(h)
	for py := int16(0); py < h; py++ {
		x1 := int16(-1)
		x2 := int16(-1)
		for _, line := range s.Lines {
			scaledY := float32(py) * hscale
			ix, intersect := getLineIntersection(0, scaledY, 255, scaledY, float32(line.A.X), float32(line.A.Y), float32(line.B.X), float32(line.B.Y))
			if intersect {
				if x1 == -1 {
					x1 = x + int16(ix/wscale)
					if x2 != -1 {
						if x1-x2 == 1 {
							x1++
						}
					}
				} else {
					x2 = x + int16(ix/wscale)
					tinydraw.LineEx(d, x1, y+py, x2, y+py, c)
					x1 = -1
				}
			}
		}
	}
}

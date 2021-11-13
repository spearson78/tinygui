package icon

import (
	"image/color"
	"math"

	"github.com/spearson78/tinygui"
	"tinygo.org/x/tinydraw"
)

//go:noinline
func Timeline(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {

	ax := int16(float32(0.12) * float32(w))
	bx := int16(float32(0.42) * float32(w))
	cx := int16(float32(0.62) * float32(w))
	dx := int16(float32(0.88) * float32(w))

	ay := int16(float32(0.67) * float32(w))
	by := int16(float32(0.38) * float32(w))
	cy := int16(float32(0.59) * float32(w))
	dy := int16(float32(0.33) * float32(w))

	r := int16(math.Max(1.0, float64(0.0085)*float64(w)))

	tinydraw.LineEx(g.Display, x+ax, y+ay, x+bx, y+by, c)
	tinydraw.LineEx(g.Display, x+bx, y+by, x+cx, y+cy, c)
	tinydraw.LineEx(g.Display, x+cx, y+cy, x+dx, y+dy, c)

	tinydraw.FilledCircleEx(g.Display, x+ax, y+ay, r, c)
	tinydraw.FilledCircleEx(g.Display, x+bx, y+by, r, c)
	tinydraw.FilledCircleEx(g.Display, x+cx, y+cy, r, c)
	tinydraw.FilledCircleEx(g.Display, x+dx, y+dy, r, c)
}

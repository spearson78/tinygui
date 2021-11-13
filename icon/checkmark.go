package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"tinygo.org/x/tinydraw"
)

func Checkmark(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {

	third := int16(float32(0.4) * float32(w))

	tinydraw.LineEx(g.Display, x, y+w-third, x+third-1, y+w, c)
	tinydraw.LineEx(g.Display, x+third, y+w, x+w-1, y, c)

	tinydraw.LineEx(g.Display, x+1, y+w-third, x+third, y+w, c)
	tinydraw.LineEx(g.Display, x+third+1, y+w, x+w, y, c)

}

package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"tinygo.org/x/tinydraw"
)

func Add(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {
	center := w / 2

	tinydraw.FilledRectangleEx(g.Display, x, y+center-1, w, 2, c)
	tinydraw.FilledRectangleEx(g.Display, x+center-1, y, 2, w, c)

}

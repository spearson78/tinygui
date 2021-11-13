package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
)

type Icon func(g *tinygui.GuiContext, x, y, w int16, c color.RGBA)

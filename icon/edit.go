package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/primitive"
)

var editShape = primitive.Shape{
	Lines: []primitive.LineSegment{

		{primitive.Point8{1, 202}, primitive.Point8{1, 254}},
		{primitive.Point8{1, 202}, primitive.Point8{157, 47}},
		{primitive.Point8{157, 47}, primitive.Point8{209, 99}},
		{primitive.Point8{209, 99}, primitive.Point8{54, 254}},

		{primitive.Point8{173, 31}, primitive.Point8{202, 3}},
		{primitive.Point8{225, 82}, primitive.Point8{173, 31}},

		{primitive.Point8{214, 3}, primitive.Point8{253, 42}},
		{primitive.Point8{253, 42}, primitive.Point8{253, 53}},
		{primitive.Point8{253, 53}, primitive.Point8{225, 82}},
	},
}

func Edit(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {
	primitive.FillShape(g.Display, x, y, w, w, &editShape, c)
}

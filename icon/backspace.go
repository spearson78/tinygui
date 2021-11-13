package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/primitive"
)

var backspaceShape = primitive.Shape{
	Lines: []primitive.LineSegment{

		//Left
		{primitive.Point8{1, 128}, primitive.Point8{64, 32}},
		{primitive.Point8{1, 128}, primitive.Point8{64, 223}},

		//Cross Left
		{primitive.Point8{110, 74}, primitive.Point8{94, 89}},
		{primitive.Point8{132, 128}, primitive.Point8{94, 89}},
		{primitive.Point8{132, 128}, primitive.Point8{94, 165}},
		{primitive.Point8{110, 181}, primitive.Point8{94, 165}},

		//Cross Middle
		{primitive.Point8{110, 74}, primitive.Point8{148, 112}},
		{primitive.Point8{186, 74}, primitive.Point8{148, 112}},

		{primitive.Point8{110, 181}, primitive.Point8{148, 143}},
		{primitive.Point8{186, 181}, primitive.Point8{148, 143}},

		//Cross Right
		{primitive.Point8{186, 74}, primitive.Point8{201, 89}},
		{primitive.Point8{163, 128}, primitive.Point8{201, 89}},
		{primitive.Point8{163, 128}, primitive.Point8{201, 165}},
		{primitive.Point8{186, 181}, primitive.Point8{201, 165}},

		//Right
		{primitive.Point8{237, 32}, primitive.Point8{254, 47}},
		{primitive.Point8{254, 47}, primitive.Point8{254, 205}},
		{primitive.Point8{254, 205}, primitive.Point8{237, 223}},
	},
}

func Backspace(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {
	primitive.FillShape(g.Display, x, y, w, w, &backspaceShape, c)
}

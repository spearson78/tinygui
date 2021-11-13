package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/primitive"
)

var calcShape = primitive.Shape{
	Lines: []primitive.LineSegment{

		//Top Left
		{primitive.Point8{1, 30}, primitive.Point8{30, 1}},
		//Middle Left
		{primitive.Point8{1, 30}, primitive.Point8{1, 225}},

		//Bottom Left
		{primitive.Point8{1, 225}, primitive.Point8{30, 254}},

		//Subtract Left
		{primitive.Point8{47, 69}, primitive.Point8{47, 91}},
		{primitive.Point8{119, 69}, primitive.Point8{119, 91}},

		//Multiply
		{primitive.Point8{145, 60}, primitive.Point8{159, 45}},
		{primitive.Point8{145, 100}, primitive.Point8{163, 80}},
		{primitive.Point8{159, 115}, primitive.Point8{145, 100}},
		{primitive.Point8{163, 80}, primitive.Point8{145, 60}},

		{primitive.Point8{159, 45}, primitive.Point8{179, 65}},
		{primitive.Point8{179, 95}, primitive.Point8{159, 115}},
		{primitive.Point8{179, 65}, primitive.Point8{199, 45}},
		{primitive.Point8{199, 115}, primitive.Point8{179, 95}},
		{primitive.Point8{195, 80}, primitive.Point8{215, 100}},
		{primitive.Point8{215, 60}, primitive.Point8{195, 80}},
		{primitive.Point8{199, 45}, primitive.Point8{215, 60}},
		{primitive.Point8{215, 100}, primitive.Point8{199, 115}},

		//Plus
		{primitive.Point8{73, 137}, primitive.Point8{73, 166}},
		{primitive.Point8{95, 137}, primitive.Point8{95, 166}},
		{primitive.Point8{44, 166}, primitive.Point8{44, 186}},
		{primitive.Point8{122, 166}, primitive.Point8{122, 186}},
		{primitive.Point8{73, 187}, primitive.Point8{73, 216}},
		{primitive.Point8{95, 187}, primitive.Point8{95, 216}},

		//Equals
		{primitive.Point8{144, 148}, primitive.Point8{144, 169}},
		{primitive.Point8{215, 148}, primitive.Point8{215, 169}},

		{primitive.Point8{144, 183}, primitive.Point8{144, 205}},
		{primitive.Point8{215, 183}, primitive.Point8{215, 205}},

		//TOP right
		{primitive.Point8{225, 1}, primitive.Point8{254, 30}},
		//MIDDLE RIGHT
		{primitive.Point8{254, 30}, primitive.Point8{254, 225}},

		//Bottom Right
		{primitive.Point8{225, 254}, primitive.Point8{254, 225}},
	},
}

func Calc(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {

	primitive.FillShape(g.Display, x, y, w, w, &calcShape, c)

}

package icon

import (
	"image/color"

	"github.com/spearson78/tinygui"
	"github.com/spearson78/tinygui/primitive"
)

var announcementShape = primitive.Shape{
	Lines: []primitive.LineSegment{

		//Top Left
		{primitive.Point8{22, 39}, primitive.Point8{38, 22}},
		//Middle Left
		{primitive.Point8{22, 39}, primitive.Point8{23, 232}},

		//Exclaim Top
		{primitive.Point8{116, 53}, primitive.Point8{116, 118}},
		{primitive.Point8{139, 53}, primitive.Point8{139, 118}},

		//Exclaim Bottom
		{primitive.Point8{116, 137}, primitive.Point8{116, 161}},
		{primitive.Point8{139, 137}, primitive.Point8{139, 161}},

		//TOP right
		{primitive.Point8{232, 39}, primitive.Point8{217, 22}},
		//MIDDLE RIGHT
		{primitive.Point8{232, 39}, primitive.Point8{232, 174}},

		//Bottom Right bubble
		{primitive.Point8{23, 232}, primitive.Point8{63, 190}},

		//Bottom Right triangle
		{primitive.Point8{232, 174}, primitive.Point8{216, 190}},
	},
}

func Announcement(g *tinygui.GuiContext, x, y, w int16, c color.RGBA) {
	primitive.FillShape(g.Display, x, y, w, w, &announcementShape, c)
}

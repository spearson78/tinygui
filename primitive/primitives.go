package primitive

import (
	"image/color"

	"github.com/spearson78/tinyamifont"
	"github.com/spearson78/tinygui"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
)

//go:noinline
func OutlineBox(d tinygui.Displayer, x, y, w, h int16, c color.RGBA) {
	tinydraw.LineEx(d, x+1, y, x+w-2, y, c)
	tinydraw.LineEx(d, x, y+1, x, y+h-2, c)
	tinydraw.LineEx(d, x+w-1, y+1, x+w-1, y+h-2, c)
	tinydraw.LineEx(d, x+1, y+h-1, x+w-2, y+h-1, c)
}

//go:noinline
func FilledBox(d tinygui.Displayer, x, y, w, h int16, c color.RGBA) {
	tinydraw.LineEx(d, x+1, y, x+w-2, y, c)
	tinydraw.FilledRectangleEx(d, x, y+1, w, h-2, c)
	tinydraw.LineEx(d, x+1, y+h-1, x+w-2, y+h-1, c)
}

//go:noinline
func OutlineBoxWithShadow(d tinygui.Displayer, x, y, w, h int16, c, s1, s2 color.RGBA) {
	OutlineBox(d, x, y, w, h-2, c)
	tinydraw.LineEx(d, x+1, y+h-2, x+w-2, y+h-2, s1)
	tinydraw.LineEx(d, x+2, y+h-1, x+w-4, y+h-1, s2)
}

//go:noinline
func FilledBoxWithShadow(d tinygui.Displayer, x, y, w, h int16, c, s1, s2 color.RGBA) {
	FilledBox(d, x, y, w, h-2, c)
	tinydraw.LineEx(d, x+1, y+h-2, x+w-2, y+h-2, s1)
	tinydraw.LineEx(d, x+2, y+h-1, x+w-4, y+h-1, s2)
}

//go:noinline
func FilledCircle(d tinygui.Displayer, x, y, w int16, c color.RGBA) {

	r := (w / 2) - 2
	tinydraw.FilledCircleEx(d, x+r+2, y+r, r, c)
}

//go:noinline
func FilledCircleWithShadow(d tinygui.Displayer, x, y, w int16, c, s1, s2 color.RGBA) {
	r := (w / 2) - 2

	tinydraw.Circle(d, x+r+2, y+r+2, r, s2)
	tinydraw.Circle(d, x+r+2, y+r+1, r, s1)
	tinydraw.FilledCircleEx(d, x+r+2, y+r, r, c)
}

//go:noinline
func WriteLine(display drivers.Displayer, font *tinyamifont.Font, x int16, y int16, str string, c color.RGBA) {
	tinyamifont.PrintString(font, display, str, x, y, c, tinyamifont.Regular)
}

//go:noinline
func LineWidth(font *tinyamifont.Font, str string) (innerWidth uint32, outboxWidth uint32) {
	w := tinyamifont.LineWidth(font, str, tinyamifont.Regular)
	return uint32(w), uint32(w)
}

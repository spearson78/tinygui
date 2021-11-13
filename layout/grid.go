package layout

import "github.com/spearson78/tinygui/component"

type GridLayout struct {
	x          int16
	y          int16
	currCell   int16
	currRow    int16
	cellWidth  int16
	cellHeight int16
	margin     int16
}

func Grid(x, y, cellWidth, cellHeight, margin int16) GridLayout {
	return GridLayout{
		x:          x,
		y:          y,
		cellWidth:  cellWidth + margin,
		cellHeight: cellHeight + margin,
	}
}

func (g *GridLayout) NextCell() component.ComponentPos {
	pos := component.ComponentPos{
		X: g.x + (g.currCell * g.cellWidth),
		Y: g.y + (g.currRow * g.cellHeight),
	}
	g.currCell++
	return pos
}

func (g *GridLayout) EndRow() component.ComponentPos {
	pos := component.ComponentPos{
		X: g.x + (g.currCell * g.cellWidth),
		Y: g.y + (g.currRow * g.cellHeight),
	}
	g.currCell = 0
	g.currRow++
	return pos
}

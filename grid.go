package main

type Grid struct {
	Rows, Cols int
	grid       [][]*Cell
}

func NewGrid(rows, cols int) *Grid {
	g := new(Grid)
	g.Rows, g.Cols = rows, cols
	g.prepareGrid()
	g.configureCells()

	return g
}

func (g *Grid) prepareGrid() {
	rows := make([][]*Cell, g.Rows)

	for r := range rows {
		cols := make([]*Cell, g.Cols)
		for c := range cols {
			cols[c] = NewCell(r, c)
		}
		rows[r] = cols
	}
	g.grid = rows
}

func (g *Grid) configureCells() {
	for r := range g.grid {
		for c := range g.grid[r] {
			g.grid[r][c].North = g.getCell(r-1, c)
			g.grid[r][c].South = g.getCell(r+1, c)
			g.grid[r][c].West = g.getCell(r, c-1)
			g.grid[r][c].East = g.getCell(r, c+1)
		}
	}
}

func (g *Grid) getCell(row, col int) *Cell {
	if row < 0 || row >= g.Rows {
		return nil
	}
	if col < 0 || col >= g.Cols {
		return nil
	}
	return g.grid[row][col]
}

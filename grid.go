package main

import (
	"math/rand"
	"time"
)

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

func (g *Grid) Size() int {
	return g.Rows * g.Cols
}

func (g *Grid) RandomCell() *Cell {
	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(g.Rows)
	col := rand.Intn(g.Cols)
	return g.grid[row][col]
}

func (g *Grid) EachRow() chan []*Cell {
	c := make(chan []*Cell)
	go func() {
		for _, r := range g.grid {
			c <- r
		}
		close(c)
	}()
	return c
}

func (g *Grid) EachCell() chan *Cell {
	c := make(chan *Cell)
	go func() {
		for r := range g.EachRow() {
			for _, i := range r {
				c <- i
			}
		}
		close(c)
	}()
	return c
}

func (g *Grid) String() string {
	output := "+"
	for i := 0; i < g.Cols; i++ {
		output += "---+"
	}
	output += "\n"

	for r := range g.EachRow() {
		top := "|"
		bottom := "+"
		for _, c := range r {
			body := "   "
			eastBoundary := ""
			if c.IsLinked(c.East) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}
			top += body
			top += eastBoundary

			southBoundary := ""
			if c.IsLinked(c.South) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}

			corner := "+"
			bottom += southBoundary
			bottom += corner
		}
		output += top + "\n"
		output += bottom + "\n"
	}

	return output
}

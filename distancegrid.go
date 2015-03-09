package main

import "strconv"

type DistanceGrid struct {
	Grid
	distances *Distances
}

func NewDistanceGrid(rows, cols int) *DistanceGrid {
	g := new(DistanceGrid)
	g.distances = new(Distances)
	g.Rows, g.Cols = rows, cols
	g.prepareGrid()
	g.configureCells()
	return g
}

func (d *DistanceGrid) SetDistances(distances *Distances) {
	d.distances = distances
}

func (g *DistanceGrid) ContentsOf(cell *Cell) string {
	return strconv.FormatInt(int64(g.distances.Distance(cell)), 36)
}

func (g *DistanceGrid) String() string {
	output := "+"
	for i := 0; i < g.Cols; i++ {
		output += "---+"
	}
	output += "\n"

	for r := range g.EachRow() {
		top := "|"
		bottom := "+"
		for _, c := range r {
			body := " " + g.ContentsOf(c) + " "
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

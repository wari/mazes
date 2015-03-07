package main

type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

func NewDistances(root *Cell) *Distances {
	d := new(Distances)
	d.root = root
	d.cells[root] = 0
	return d
}

func (d *Distances) Cells() []*Cell {
	cells := make([]*Cell, len(d.cells))
	for k, v := range d.cells {
		cells[v] = k
	}
	return cells
}

func (d *Distances) Distance(cell *Cell) int {
	return d.cells[cell]
}

func (d *Distances) SetDistance(cell *Cell, distance int) {
	d.cells[cell] = distance
}

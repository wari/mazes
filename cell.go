package main

type Cell struct {
	row, col                 int
	North, South, East, West *Cell
	links                    map[*Cell]bool
}

func NewCell(row, col int) *Cell {
	c := new(Cell)
	c.row, c.col = row, col
	c.links = make(map[*Cell]bool)
	return c
}

func (c Cell) Row() int {
	return c.row
}

func (c Cell) Col() int {
	return c.col
}

func (c *Cell) Link(cell *Cell, bidi bool) *Cell {
	c.links[cell] = true

	if bidi {
		cell.Link(c, false)
	}
	return c
}

func (c *Cell) Unlink(cell *Cell, bidi bool) *Cell {
	delete(c.links, cell)
	if bidi {
		cell.Unlink(c, false)
	}
	return c
}

func (c Cell) IsLinked(cell *Cell) bool {
	for x := range c.links {
		if x == cell {
			return true
		}
	}
	return false
}

func (c Cell) Links() []*Cell {
	list := make([]*Cell, 0)
	for x := range c.links {
		list = append(list, x)
	}
	return list
}

func (c Cell) Neighbors() []*Cell {
	list := make([]*Cell, 0)
	if c.North != nil {
		list = append(list, c.North)
	}
	if c.South != nil {
		list = append(list, c.South)
	}
	if c.East != nil {
		list = append(list, c.East)
	}
	if c.West != nil {
		list = append(list, c.West)
	}
	return list
}

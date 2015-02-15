// Not done yet, currently, what's below is just test code to see what's going on
package main

import "fmt"

func main() {
	c := NewCell(0, 0)
	d := NewCell(0, 1)
	c.Link(d, true)
	c.North = d
	d.South = c
	d.Link(c, false)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(c.Links())
	fmt.Println(c.Neighbors())
	fmt.Println(c.IsLinked(d))
	c.Unlink(d, true)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(c.Links())
	fmt.Println(c.Neighbors())
	fmt.Println(c.IsLinked(d))

	g := NewGrid(10, 10)
	fmt.Println(g.grid)

	fmt.Println(g.grid[4][5])
	fmt.Println(g.grid[4][5].North)
	fmt.Println(g.grid[4][5].South)
	fmt.Println(g.grid[4][5].East)
	fmt.Println(g.grid[4][5].West)

	for y := range g.EachRow() {
		fmt.Println(y)
	}
}

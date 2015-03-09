// Not done yet, currently, what's below is just test code to see what's going on
package main

import "fmt"

func main() {
	fmt.Println("Binary Tree")
	g := NewDistanceGrid(10, 10)
	b := new(BinaryTree)
	b.On(g)
	start := g.GetCell(0, 0)
	distances := start.Distances()
	g.SetDistances(distances)
	fmt.Println(g)
	fmt.Println("Sidewinder")
	g = NewDistanceGrid(10, 10)
	s := new(Sidewinder)
	s.On(g)
	start = g.GetCell(0, 0)
	distances = start.Distances()
	g.SetDistances(distances)
	fmt.Println(g)
	fmt.Println("Showing image")
}

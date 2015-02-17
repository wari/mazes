// Not done yet, currently, what's below is just test code to see what's going on
package main

import "fmt"

func main() {
	fmt.Println("Binary Tree")
	g := NewGrid(10, 10)
	b := new(BinaryTree)
	b.On(g)
	fmt.Println(g)
	fmt.Println("Sidewinder")
	g = NewGrid(10, 10)
	s := new(Sidewinder)
	s.On(g)
	fmt.Println(g)
}

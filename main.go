// Not done yet, currently, what's below is just test code to see what's going on
package main

import "fmt"

func main() {
	g := NewGrid(20, 20)
	b := new(BinaryTree)
	b.On(g)
	fmt.Println(g)
}

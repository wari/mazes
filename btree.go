package main

import (
	"math/rand"
	"time"
)

type BinaryTree struct{}

func (b *BinaryTree) On(g *Grid) {
	rand.Seed(time.Now().UnixNano())
	for c := range g.EachCell() {
		neighbors := make([]*Cell, 0)
		if c.South != nil {
			neighbors = append(neighbors, c.South)
		}
		if c.East != nil {
			neighbors = append(neighbors, c.East)
		}
		l := len(neighbors)
		if l != 0 {
			i := rand.Intn(len(neighbors))
			neighbor := neighbors[i]
			if neighbor != nil {
				c.Link(neighbor, true)
			}
		}
	}
}

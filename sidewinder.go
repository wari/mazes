package main

import (
	"math/rand"
	"time"
)

type Sidewinder struct{}

type cells []*Cell

func (r cells) sample() *Cell {
	l := len(r)
	if l == 0 {
		return nil
	} else {
		return r[rand.Intn(l)]
	}
}

func (s *Sidewinder) On(g Grider) Grider {
	rand.Seed(time.Now().UnixNano())
	for r := range g.EachRow() {
		run := make(cells, 0)
		for _, c := range r {
			run = append(run, c)

			atEasternBoundary := c.East == nil
			atNorthernBoundary := c.North == nil
			shouldCloseOut := atEasternBoundary || (!atNorthernBoundary && rand.Intn(2) == 0)

			if shouldCloseOut {
				member := run.sample()
				if member.North != nil {
					member.Link(member.North, true)
				}
				run = nil // Clear out the cells
			} else {
				c.Link(c.East, true)
			}
		}
	}
	return g
}

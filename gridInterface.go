package main

type Grider interface {
	EachRow() chan []*Cell
	EachCell() chan *Cell
}

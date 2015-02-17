package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/exec"
)

var (
	white color.Color = color.RGBA{255, 255, 255, 255}
	black color.Color = color.RGBA{0, 0, 0, 255}
)

type Image struct {
	size int
}

func NewImage(size int) *Image {
	return &Image{size: size}
}

func (i *Image) DrawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	for i := x1; i < x2; i++ {
		img.Set(i, y2, c)
	}
	for i := y1; i < y2; i++ {
		img.Set(x2, i, c)
	}
}

func (i *Image) DrawGrid(g *Grid) *image.RGBA {
	width := i.size * g.Cols
	height := i.size * g.Rows
	background := white
	wall := black

	img := image.NewRGBA(image.Rect(0, 0, width+1, height+1))
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	for c := range g.EachCell() {
		x1 := c.Col() * i.size
		y1 := c.Row() * i.size
		x2 := (c.Col() + 1) * i.size
		y2 := (c.Row() + 1) * i.size

		if c.North == nil {
			i.DrawLine(img, x1, y1, x2, y1, wall)
		}
		if c.West == nil {
			i.DrawLine(img, x1, y1, x1, y2, wall)
		}
		if !c.IsLinked(c.East) {
			i.DrawLine(img, x2, y1, x2, y2, wall)
		}
		if !c.IsLinked(c.South) {
			i.DrawLine(img, x1, y2, x2, y2, wall)
		}
	}

	return img
}

func (i *Image) Save(filename string, img *image.RGBA) {
	w, _ := os.Create(filename)
	defer w.Close()
	png.Encode(w, img) //Encode writes the Image m to w in PNG format.
}

func Show(name string) {
	command := "xdg-open"
	cmd := exec.Command(command, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

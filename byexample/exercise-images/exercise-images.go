package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

// At implements image.Image.
func (i Image) At(x int, y int) color.Color {
	v := uint8(0)
	return color.RGBA{v, v, 255, 255}
}

// Bounds implements image.Image.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// ColorModel implements image.Image.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}

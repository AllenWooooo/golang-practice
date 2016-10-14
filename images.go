package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) At(x, y int) color.Color {
	r := uint8(x ^ y)
	g := uint8(x*x + y*y)
	b := uint8(x * y)
	a := uint8(math.Sin(float64(x*y)) * 255)

	return color.RGBA{r, g, b, a}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

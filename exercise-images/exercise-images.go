package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 200)
}

func (i Image) At(x, y int) color.Color {
	val := uint8((x + y) / 2)
	return color.RGBA{val, val, 255, 255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

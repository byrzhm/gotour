package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.GrayModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.Gray{uint8(x ^ y)}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}

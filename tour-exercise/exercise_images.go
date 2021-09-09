package tourexercise

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// https://tour.golang.org/methods/25
// Exercise: Images

type myImage struct {
	w, h int
}

func (img *myImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *myImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *myImage) At(x, y int) color.Color {
	v := uint8((x ^ y) * (x ^ y))
	return color.RGBA{v, v, 255, 255}
}

// Solution of "images" exercise
//
// For detail:
// https://tour.golang.org/methods/25
//
// For visual result:
// https://play.golang.org/p/_4_JwNUlP60
func Images() {
	m := &myImage{200, 200}
	pic.ShowImage(m)
}

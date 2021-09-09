package methods

import (
	"fmt"
)

// https://tour.golang.org/methods/25
// Exercise: Images

// Remember the picture generator you wrote earlier?
// (tour\moretypes\mt18_exercise_slices.go)
// Let's write another one, but this time it will
// return an implementation of image.Image
// instead of a slice of data.
//
// Define your own Image type, implement
// the necessary methods, and call pic.ShowImage.
//
// Bounds should return a image.Rectangle,
// like image.Rect(0, 0, w, h).
//
// ColorModel should return color.RGBAModel.
//
// At should return a color; the value v
// in the last picture generator corresponds
// to color.RGBA{v, v, 255, 255} in this one.
func ExImages() {
	fmt.Println("Please read and implement :)")
}

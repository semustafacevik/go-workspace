package methods

import (
	"fmt"
	"math"
)

// https://tour.golang.org/methods/2
// Methods are functions

func abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Remember: a method is just a function
// with a receiver argument.
//
// Here's Abs written as a regular function
// with no change in functionality.
func MethodsFunc() {
	v := vertex{5, 12}
	fmt.Println(abs(v))
}

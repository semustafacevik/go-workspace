package tour

import (
	"fmt"
	"math"
)

// https://tour.golang.org/basics/13
// Type conversions

// The expression T(v) converts the value v to the type T.
//
// Some numeric conversions:
// 	var i int = 42
// 	var f float64 = float64(i)
// 	var u uint = uint(f)
//
// Or put more simply:
// 	i := 42
// 	f := float64(i)
// 	u := uint(f)
//
// Unlike in C, in Go assignment between items of different type
// requires an explicit conversion. Try removing the float64
// or uint conversions in the sexample and see what happens.
func TypeConversions() {
	// i := 42
	// var u uint = i // -> var u uint = uint(i)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(x, y, f)
}

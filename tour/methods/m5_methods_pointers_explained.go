package methods

import "fmt"

// https://tour.golang.org/methods/5
// Pointers and functions

func scale(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here we see the Scale method
// rewritten as functions.
//
// Again, try removing the * from line 8.
// Can you see why the behavior changes?
// What else did you need to change for
// the example to compile?
//
// (If you're not sure,
// continue to the next page.)
func MethodsPointersExplained() {
	v := vertex{5, 12}
	scale(&v, 10)
	fmt.Println(abs(v))
}

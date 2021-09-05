package methods

import (
	"fmt"
	"math"
)

// https://tour.golang.org/methods/8
// Choosing a value or pointer receiver

func (v *vertex) ScaleM8(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *vertex) AbsM8() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// There are two reasons to use
// a pointer receiver.
//
// The first is so that the method
// can modify the value that its
// receiver points to.
//
// The second is to avoid copying
// the value on each method call.
// This can be more efficient if the receiver
// is a large struct, for example.
//
// In this example, both Scale and Abs are
// with receiver type *Vertex, even though
// the Abs method needn't modify its receiver.
//
// In general, all methods on a given type
// should have either value or
// pointer receivers but not a mixture of both.
// (We'll see why over the next few pages.)
func MethodsWithPointerReceivers() {
	v := &vertex{3, 4}
	fmt.Printf("Before scalling: %+v - Abs: %v\n", v, v.AbsM8())
	v.ScaleM8(2)
	fmt.Printf("After scalling : %+v - Abs: %v\n", v, v.AbsM8())
}

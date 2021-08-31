package methods

import "fmt"

// https://tour.golang.org/methods/9
// Interfaces

type Abser interface {
	AbsM9() float64
}

// An interface type is defined as a set
// of method signatures.
//
// A value of interface type can hold
// any value that implements those methods.
//
// Note: There is an error in the example
// code on line 37. Vertex (the value type)
// doesn't implement Abser because
// the Abs method is defined only on *Vertex
// (the pointer type).
func Interfaces() {
	var a Abser

	f := MyFloat(-100.89)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.AbsM9())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.AbsM9())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v
}

func (f MyFloat) AbsM9() float64 {
	return f.Abs()
}

func (v *Vertex) AbsM9() float64 {
	return v.Abs()
}

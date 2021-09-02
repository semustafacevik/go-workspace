package methods

import "fmt"

// https://tour.golang.org/methods/6
// Methods and pointer indirection

// defined in m4_methods_pointers
// func (v *Vertex) Scale(f float64) {
// 	v.X = v.X * f
// 	v.Y = v.Y * f
// }

func ScaleFunc(v *Vertex, f float64) {
	Scale(v, f)
}

// Comparing the previous two programs,
// you might notice that functions with
// a pointer argument must take a pointer:
//
//  var v Vertex
//  ScaleFunc(v, 5)  // Compile error!
//  ScaleFunc(&v, 5) // OK
//
// while methods with pointer receivers
// take either a value or a pointer as
// the receiver when they are called:
//
//  var v Vertex
//  v.Scale(5)  // OK
//  p := &v
//  p.Scale(10) // OK
//
// For the statement v.Scale(5), even though
// v is a value and not a pointer, the method
// with the pointer receiver is called
// automatically. That is, as a convenience,
// Go interprets the statement v.Scale(5) as
// (&v).Scale(5) since the Scale method has
// a pointer receiver.
func Indirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 5)

	p := &Vertex{12, 5}
	p.Scale(3)
	ScaleFunc(p, 7)

	fmt.Println(v, p)
}

package methods

import "fmt"

// https://tour.golang.org/methods/7
// Methods and pointer indirection (2)

// defined in m1_methods
// func (v Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

func absFunc(v vertex) float64 {
	return abs(v)
}

// The equivalent thing happens
// in the reverse direction.
//
// Functions that take a value argument
// must take a value of that specific type:
//
//  var v Vertex
//  fmt.Println(AbsFunc(v))  // OK
//  fmt.Println(AbsFunc(&v)) // Compile error!
//
// while methods with value receivers
// take either a value or a pointer as
// the receiver when they are called:
//
//  var v Vertex
//  fmt.Println(v.Abs()) // OK
//  p := &v
//  fmt.Println(p.Abs()) // OK
//
// In this case, the method call p.Abs()
// is interpreted as (*p).Abs().
func IndirectionValues() {
	v := vertex{3, 4}
	fmt.Println(v.Abs(), absFunc(v))

	p := &vertex{5, 12}
	fmt.Println(p.Abs(), absFunc(*p))
}

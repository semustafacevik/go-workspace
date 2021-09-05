package methods

import (
	"fmt"
)

// https://tour.golang.org/methods/3
// Methods continued

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can declare a method on
// non-struct types, too.
//
// In this example we see a numeric type
// MyFloat with an Abs method.
//
// You can only declare a method with
// a receiver whose type is defined
// in the same package as the method.
// You cannot declare a method with a
// receiver whose type is defined in
// another package (which includes
// the built-in types such as int).
func MethodsContunied() {
	f := myFloat(-100.89)
	fmt.Println(f.Abs())
}

package methods

import (
	"fmt"
)

// https://tour.golang.org/methods/16
// Type switches

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case bool:
		fmt.Printf("The opposite of %v is %v\n", v, !v)
	default:
		fmt.Printf("Unknown type -> %T\n", v)
	}
}

// A type switch is a construct that
// permits several type assertions in series.
//
// A type switch is like a regular
// switch statement, but the cases in
// a type switch specify types (not values),
// and those values are compared against
// the type of the value held by
// the given interface value.
//
//  switch v := i.(type) {
//  case T:
//     // here v has type T
//  case S:
//     // here v has type S
//  default:
//     // no match; here v has the same type as i
//  }
//
// The declaration in a type switch has
// the same syntax as a type assertion i.(T),
// but the specific type T is replaced
// with the keyword type.
//
// This switch statement tests whether
// the interface value i holds a value of
// type T or S. In each of the T and S cases,
// the variable v will be of type T or S
// respectively and hold the value held by i.
// In the default case (where there is no match),
// the variable v is of the same interface type
// and value as i.
func TypeSwitches() {
	do(4)
	do("go-workspace")
	do(false)
	do(4.01)
}

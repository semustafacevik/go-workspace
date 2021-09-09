package methods

import "fmt"

// https://tour.golang.org/methods/11
// Interface values

type f float64

func (f f) M() {
	fmt.Println(f)
}

// Under the hood, interface values can be
// thought of as a tuple of a value and
// a concrete type:
//
//  (value, type)
//
// An interface value holds a value of
// a specific underlying concrete type.
//
// Calling a method on an interface value
// executes the method of the same name
// on its underlying type.
func InterfaceValues() {
	var i myInterface

	i = &myStruct{"go-workspace"}
	describe(i)
	i.M()

	fmt.Println("   -----")

	i = f(-100.89)
	describe(i)
	i.M()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

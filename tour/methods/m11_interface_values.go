package methods

import "fmt"

// https://tour.golang.org/methods/11
// Interface values

type F float64

func (f F) M() {
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
	var i I

	i = &T{"go-workspace"}
	describe(i)
	i.M()

	fmt.Println("   -----")

	i = F(-100.89)
	describe(i)
	i.M()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

package methods

import "fmt"

// https://tour.golang.org/methods/12
// Interface values with nil underlying values

type myInterface2 interface {
	MWithNilCheck()
}

func (st *myStruct) MWithNilCheck() {
	if st == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(st.S)
}

// If the concrete value inside the interface
// itself is nil, the method will be called
// with a nil receiver.
//
// In some languages this would trigger
// a null pointer exception, but in Go it is
// common to write methods that gracefully
// handle being called with a nil receiver
// (as with the method M in this example.)
//
// Note that an interface value that holds
// a nil concrete value is itself non-nil.
func InterfaceValuesWithNil() {
	var i myInterface2
	var st *myStruct

	i = st
	describe(i)
	i.MWithNilCheck()

	fmt.Println("   -----")

	i = &myStruct{"go-workspace"}
	describe(i)
	i.MWithNilCheck()
}

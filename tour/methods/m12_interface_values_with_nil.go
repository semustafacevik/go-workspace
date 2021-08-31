package methods

import "fmt"

// https://tour.golang.org/methods/12
// Interface values with nil underlying values

type I2 interface {
	MWithNilCheck()
}

func (t *T) MWithNilCheck() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
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
	var i I2
	var t *T

	i = t
	describe(i)
	i.MWithNilCheck()

	fmt.Println("   -----")

	i = &T{"hello"}
	describe(i)
	i.MWithNilCheck()
}

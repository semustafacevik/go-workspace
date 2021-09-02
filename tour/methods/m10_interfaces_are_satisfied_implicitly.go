package methods

import "fmt"

// https://tour.golang.org/methods/10
// Interfaces are implemented implicitly

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// A type implements an interface by
// implementing its methods. There is
// no explicit declaration of intent,
// no "implements" keyword.
//
// Implicit interfaces decouple the definition
// of an interface from its implementation,
// which could then appear in any package
// without prearrangement.
func InterfacesAreSatisfiedImplicitly() {
	var i I = T{"go-workspace"}
	i.M()
}

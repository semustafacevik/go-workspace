package methods

import "fmt"

// https://tour.golang.org/methods/10
// Interfaces are implemented implicitly

type myInterface interface {
	M()
}

type myStruct struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (st myStruct) M() {
	fmt.Println(st.S)
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
	var i myInterface = myStruct{"go-workspace"}
	i.M()
}

package tour

import "fmt"

var i, j int = -1, 0

// https://tour.golang.org/basics/9
// Variables with initializers

// A var declaration can include
// initializers, one per variable.
//
// If an initializer is present,
// the type can be omitted; the variable
// will take the type of the initializer.
func VariablesWithInitializers() {
	var c, python, java, golang = true, false, 1, "studying:)"
	fmt.Println(i, j, c, python, java, golang)
}

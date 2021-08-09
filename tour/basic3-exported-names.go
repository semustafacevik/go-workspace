package tour

import (
	"fmt"
	"math"
)

/*
	https://tour.golang.org/basics/3

	Exported names
	In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

	pizza and pi do not start with a capital letter, so they are not exported.

	When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

	Run the code. Notice the error message.

	To fix the error, rename math.pi to math.Pi and try it again.
*/
func ExportedNames() {
	// fmt.Println(math.pi) cannot refer to unexported name math.pi

	fmt.Println(math.Pi)
}

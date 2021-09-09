package basics

import (
	"fmt"
	"math"
)

// https://tour.golang.org/basics/2
// Imports

// This code groups the imports into a parenthesized,
// "factored" import statement.
//
// You can also write multiple import statements, like:
// 	import "fmt"
// 	import "math"
//
// But it is good style to use the factored import statement.
func Imports() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

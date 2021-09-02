package moretypes

import "fmt"

// https://tour.golang.org/moretypes/6
// Arrays

// The type [n]T is an array of n values
// of type T.
//
// The expression
//
//  var a [10]int
// declares a variable a as an array
// of ten integers.
//
// An array's length is part of its type,
// so arrays cannot be resized. This seems
// limiting, but don't worry; Go provides
// a convenient way of working with arrays.
func Arrays() {
	var a [3]string
	a[0] = "go"
	a[1] = "workspace"
	a[2] = ";)"

	fmt.Println(a[0], a[1], a[2])
	fmt.Println(a)

	fmt.Println()

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

package tour

import "fmt"

// https://tour.golang.org/moretypes/10
// Slice defaults

// When slicing, you may omit the high
//or low bounds to use their defaults
// instead.
//
// The default is zero for the low bound
// and the length of the slice for
// the high bound.
//
// For the array
//
//  var a [10]int
//
// these slice expressions are equivalent:
//
//  a[0:10]
//  a[:10]
//  a[0:]
//  a[:]
//
func SliceBounds() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	primes = primes[1:6] // primes -> [3,5,7,11,13]
	fmt.Println(primes)

	primes = primes[:4] // primes -> [3,5,7,11]
	fmt.Println(primes)

	primes = primes[1:] // primes -> [5,7,11]
	fmt.Println(primes)
}

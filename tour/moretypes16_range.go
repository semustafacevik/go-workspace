package tour

import "fmt"

// https://tour.golang.org/moretypes/16
// Range

var powersOfTwo = []int{1, 2, 4, 8, 16, 32, 64, 128}

// The range form of the for loop
// iterates over a slice or map.
//
// When ranging over a slice, two values
// are returned for each iteration.
// The first is the index, and the second
// is a copy of the element at that index.
func Range() {
	for i, v := range powersOfTwo {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}

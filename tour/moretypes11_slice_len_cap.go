package tour

import "fmt"

// https://tour.golang.org/moretypes/11
// Slice length and capacity

// A slice has both a length and a capacity.
//
// The length of a slice is the number of
// elements it contains.
//
// The capacity of a slice is the number of
// elements in the underlying array, counting
// from the first element in the slice.
//
// The length and capacity of a slice s can
// be obtained using the expressions len(s)
// and cap(s).
//
// You can extend a slice's length by
// re-slicing it, provided it has sufficient
// capacity. Try changing one of the slice
// operations in the example program to extend
// it beyond its capacity and see what happens.
func SliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0] // s -> [] / cap=6
	printSlice(s)

	// Extend its length.
	s = s[:4] // s -> [2,3,5,7] / cap=6
	printSlice(s)

	// Drop its first two values.
	s = s[2:] // s -> [5,7] / cap=4
	printSlice(s)

	s = s[:4] // s -> [5,7,11,13] / cap=4
	printSlice(s)

	s = s[1:3] // s -> [7,11] / cap=3
	printSlice(s)

	// s = s[:4] // slice bounds out of range
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

package moretypes

import "fmt"

// https://tour.golang.org/moretypes/13
// Creating a slice with make

// Slices can be created with the built-in
// make function; this is how you create
// dynamically-sized arrays.
//
// The make function allocates a zeroed
// array and returns a slice that refers
// to that array:
//
// 	a := make([]int, 5)  // len(a)=5
//
// To specify a capacity, pass a third
// argument to make:
//
//  b := make([]int, 0, 5) // len(b)=0, cap(b)=5
//
//  b = b[:cap(b)] // len(b)=5, cap(b)=5
//  b = b[1:]      // len(b)=4, cap(b)=4
//
func MakingSlices() {
	a := make([]int, 5)
	printSliceWithName("a", a)

	b := make([]int, 0, 5)
	printSliceWithName("b", b)

	c := b[:2]
	printSliceWithName("c", c)

	d := c[2:5]
	printSliceWithName("d", d)
}

func printSliceWithName(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

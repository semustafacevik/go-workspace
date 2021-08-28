package tourexercise

import "fmt"

// https://tour.golang.org/moretypes/23
// Exercise: Maps

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		f := n1
		n1, n2 = n2, f+n2
		return f
	}
}

// Solution of "fibonacci closure" exercise
//
// For detail:
// https://tour.golang.org/moretypes/26
func FibonacciClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d   ", f())
	}
}

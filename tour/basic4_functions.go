package tour

import "fmt"

func add(x int, y int) int {
	return x + y
}

// https://tour.golang.org/basics/4
// Functions

// A function can take zero or more arguments.
// In this example, add takes two parameters of type int.
// Notice that the type comes after the variable name.
// (For more about why types look the way they do,
// see the article on Go's declaration syntax.)
func Functions() {
	fmt.Println(add(4, 8))
}

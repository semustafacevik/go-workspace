package basics

import "fmt"

// https://tour.golang.org/basics/6
// Multiple results

func swap(x, y string) (string, string) {
	return y, x
}

// A function can return any number of results.
//
// The swap function returns two strings.
func MultipleResult() {
	a, b := swap("go", "workspace")
	fmt.Println(a, b)
}

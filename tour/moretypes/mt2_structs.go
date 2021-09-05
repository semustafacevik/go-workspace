package moretypes

import "fmt"

// https://tour.golang.org/moretypes/2
// Structs

type vertex struct {
	X int
	Y int
}

// A struct is a collection of fields.
func Structs() {
	fmt.Println(vertex{1, 2})
}

package tour

import "fmt"

// https://tour.golang.org/moretypes/2
// Structs

type Vertex struct {
	X int
	Y int
}

// A struct is a collection of fields.
func Structs() {
	fmt.Println(Vertex{1, 2})
}

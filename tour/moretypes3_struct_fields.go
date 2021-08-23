package tour

import "fmt"

// https://tour.golang.org/moretypes/3
// Struct Fields

// Struct fields are accessed using a dot.
func StructFields() {
	v := Vertex{1, 2}
	v.X = 11

	fmt.Println(v.X, v.Y)
}

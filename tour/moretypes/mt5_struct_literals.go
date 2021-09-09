package moretypes

import "fmt"

// https://tour.golang.org/moretypes/5
// Struct Literals

var (
	v1 = vertex{1, 2}    // has type Vertex
	v2 = vertex{Y: 1}    // X:0 is implicit
	v3 = vertex{}        // X:0 and Y:0
	p  = &vertex{11, 22} // has type *Vertex
)

// A struct literal denotes a newly
// allocated struct value by listing
// the values of its fields.
//
// You can list just a subset of fields
// by using the Name: syntax.(And the order
// of named fields is irrelevant.)
//
// The special prefix & returns a pointer
// to the struct value.
func StructLiterals() {
	fmt.Println(v1, v2, v3, p, *p)
}

package moretypes

import "fmt"

// https://tour.golang.org/moretypes/21
// Map literals continued

var m3 = map[string]Location{
	"Clock Tower": {38.41800, 27.12389},
	"Google":      {37.42202, -122.08408},
}

// If the top-level type is just
// a type name, you can omit it
// from the elements of the literal.
func MapLiteralsContinued() {
	fmt.Println(m3)
}

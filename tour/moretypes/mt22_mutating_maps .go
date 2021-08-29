package moretypes

import "fmt"

// https://tour.golang.org/moretypes/22
// Mutating Maps

// Insert or update an element in map m:
//
//  m[key] = elem
//
// Retrieve an element:
//
//  elem = m[key]
//
// Delete an element:
//
//  delete(m, key)
//
// Test that a key is present with
// a two-value assignment:
//
//  elem, ok = m[key]
//
// If key is in m, ok is true. If not,
// ok is false.
//
// If key is not in the map, then elem
// is the zero value for the map's
// element type.
//
// Note: If elem or ok have not yet
// been declared you could use
// a short declaration form:
//
//  elem, ok := m[key]
//
func MutatingMaps() {
	m := make(map[string]int)

	m["Number"] = 4
	fmt.Println("The number :", m["Number"])

	m["Number"] = 8
	fmt.Println("The number :", m["Number"])

	delete(m, "Number")
	fmt.Println("The number :", m["Number"])

	v, ok := m["Number"]
	fmt.Println("The number :", v, "- Present?", ok)

}

package moretypes

import "fmt"

// https://tour.golang.org/moretypes/19
// Maps

type Location struct {
	Lat, Long float64
}

var m map[string]Location

// A map maps keys to values.
//
// The zero value of a map is nil. A nil
// map has no keys, nor can keys be added.
//
// The make function returns a map of
// the given type, initialized and ready
// for use.
func Maps() {
	m = make(map[string]Location)

	m["Clock Tower"] = Location{
		38.41800, 27.12389,
	}

	fmt.Println(m["Clock Tower"])
}

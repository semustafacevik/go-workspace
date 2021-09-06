package moretypes

import "fmt"

// https://tour.golang.org/moretypes/20
// Map literals

var m2 = map[string]location{
	// please review the next section
	"Clock Tower": location{
		38.41800, 27.12389,
	},
	"Google": location{
		37.42202, -122.08408,
	},
}

// Map literals are like struct literals,
// but the keys are required.
func MapLiterals() {
	fmt.Println(m2)
}

package tour

import "fmt"

// https://tour.golang.org/moretypes/20
// Map literals

var m2 = map[string]Location{
	// please review the next section
	"Clock Tower": Location{
		38.41800, 27.12389,
	},
	"Google": Location{
		37.42202, -122.08408,
	},
}

// Map literals are like struct literals,
// but the keys are required.
func MapLiterals() {
	fmt.Println(m2)
}

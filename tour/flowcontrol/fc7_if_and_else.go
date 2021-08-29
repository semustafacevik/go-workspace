package flowcontrol

import (
	"fmt"
	"math"
)

func powWithElse(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// https://tour.golang.org/flowcontrol/7
// If and else

// Variables declared inside an if short statement
// are also available inside any of the else blocks.
//
// (Both calls to pow return their results before
// the call to fmt.Println in main begins.)
func IfAndElse() {
	fmt.Println(
		powWithElse(3, 2, 10),
		powWithElse(3, 3, 20),
	)
}

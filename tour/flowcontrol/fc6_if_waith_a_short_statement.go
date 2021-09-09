package flowcontrol

import (
	"fmt"
	"math"
)

// https://tour.golang.org/flowcontrol/6
// If with a short statement

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Like for, the if statement can start with a short
// statement to execute before the condition.
//
// Variables declared by the statement are only in
// scope until the end of the if.
//
// (Try using v in the last return statement.)
func IfWithAShortStatement() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

package tour

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// https://tour.golang.org/flowcontrol/5
// If

// Go's if statements are like its for loops;
// the expression need not be surrounded by
// parentheses ( ) but the braces { } are required.
func If() {
	fmt.Println(sqrt(2), sqrt(-4))
}

package tourexercise

import (
	"fmt"
)

// https://tour.golang.org/flowcontrol/8
// Exercise: Loops and Functions

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v. iteration -> %v\n", i+1, z)
	}
	return z
}

// Solution of "loops and functions" exercise
//
// For detail:
// https://tour.golang.org/flowcontrol/8
func LoopsAndFunctions() {
	fmt.Println("Result ->", Sqrt(2))
}

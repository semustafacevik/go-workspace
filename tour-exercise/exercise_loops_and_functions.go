package tourexercise

import (
	"fmt"
	"math"
)

// https://tour.golang.org/flowcontrol/8
// Exercise: Loops and Functions

const diff = 1e-4

func sqrt(x float64) float64 {
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
	fmt.Println("\nResult ->", sqrt(2))
}

func sqrtV2(x float64) float64 {
	z := float64(1)
	i := 1
	var f float64

	for {
		z, f = z-(z*z-x)/(2*z), z

		// you can change the diff value to get closer results
		// e.g 1e-10
		if math.Abs(f-z) < diff {
			fmt.Printf("Found the result in the %v. iteration\n\n", i)
			break
		}
		i++
	}
	return z
}

// Solution (v2) of "loops and functions" exercise
//
// For detail:
// https://tour.golang.org/flowcontrol/8
func LoopsAndFunctionsV2() {
	myResult := sqrtV2(2)
	mathResult := math.Sqrt(2)

	fmt.Println("  my sqrt func result ->", myResult)
	fmt.Println("math.Sqrt func result ->", mathResult)
	fmt.Println("           difference ->", math.Abs(myResult-mathResult))
}

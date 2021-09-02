package tourexercise

import (
	"fmt"
	"math"
)

// https://tour.golang.org/methods/20
// Exercise: Errors

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrtWithErr(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	}

	return math.Sqrt(x), nil
}

// Solution of "errors" exercise
//
// For detail:
// https://tour.golang.org/methods/20
func Errors() {
	fmt.Println(sqrtWithErr(2))
	fmt.Println(sqrtWithErr(-2))

	fmt.Println(errNegativeSqrt(-22.3).Error())
}

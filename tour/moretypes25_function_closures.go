package tour

import "fmt"

// https://tour.golang.org/moretypes/25
// Function closures

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// Go functions may be closures. A closure
// is a function value that references
// variables from outside its body.
// The function may access and assign to
// the referenced variables; in this sense
// the function is "bound" to the variables.
//
// For example, the adder function returns
// a closure. Each closure is bound to
// its own sum variable.
func FunctionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println("-----")
	fmt.Println(pos(-45), neg(90))

}

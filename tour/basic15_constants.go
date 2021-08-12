package tour

import "fmt"

const Pi = 3.14

// https://tour.golang.org/basics/15
// Constants

// Constants are declared like variables,
// but with the const keyword.
//
// Constants can be character, string,
// boolean or numeric values.
//
// Constants cannot bedeclared using
// the := syntax.
func Constants() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

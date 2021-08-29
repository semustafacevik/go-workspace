package basics

import "fmt"

func addContinued(x, y int) int {
	return x + y
}

// https://tour.golang.org/basics/5
// Functions continued

// When two or more consecutive named function
// parameters share a type, you can omit
// the type from all but the last.
//
// In this example, we shortened
//  x int, y int  ->  x, y int
func FunctionsContinued() {
	fmt.Println(addContinued(22, 44))
}

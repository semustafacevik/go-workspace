package tour

import "fmt"

// https://tour.golang.org/flowcontrol12
// Defer

// A defer statement defers the execution
// of a function until the surrounding
// function returns.
//
// The deferred call's arguments are
// evaluated immediately, but the function
// call is not executed until the
// surrounding function returns.
func Defer() {
	defer fmt.Println("workspace")

	fmt.Println("go")
}

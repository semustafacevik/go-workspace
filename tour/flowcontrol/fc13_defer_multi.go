package flowcontrol

import "fmt"

// https://tour.golang.org/flowcontrol13
// Stacking defers

// Deferred function calls are pushed onto
// a stack. When a function returns,
// its deferred calls are executed in
// last-in-first-out order.
//
// To learn more about defer statements
// read this blog post. ->
// https://blog.golang.org/defer-panic-and-recover
func DeferMulti() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

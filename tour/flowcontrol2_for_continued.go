package tour

import "fmt"

// https://tour.golang.org/flowcontrol/2
// For continued

// The init and post statements are optional.
func ForContinued() {
	sum := 1
	for sum < 1000 { //		for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

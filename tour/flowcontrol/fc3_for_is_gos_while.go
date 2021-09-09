package flowcontrol

import "fmt"

// https://tour.golang.org/flowcontrol/3
// For is Go's "while"

// At that point you can drop the semicolons:
// C's while is spelled for in Go.
func ForIsGosWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

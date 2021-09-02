package flowcontrol

import "fmt"

// https://tour.golang.org/flowcontrol/4
// Forever

// If you omit the loop condition it loops forever,
// so an infinite loop is compactly expressed.
func Forever() {
	for {
		fmt.Println("infinite loop")
	}
}

package concurrency

import (
	"fmt"
	"time"
)

// https://tour.golang.org/concurrency/6
// Default Selection

// The default case in a select is run
// if no other case is ready.
//
// Use a default case to try a send
// or receive without blocking:
//
//  select {
//  case i := <-c:
//      // use i
//  default:
//      // receiving from c would block
//  }
func DefaultSelection() {
	tick := time.Tick(400 * time.Millisecond)
	boom := time.After(2000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

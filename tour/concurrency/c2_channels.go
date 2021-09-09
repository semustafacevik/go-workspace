package concurrency

import "fmt"

// https://tour.golang.org/concurrency/2
// Channels

func sum(n []int, c chan int) {
	sum := 0
	for _, v := range n {
		sum += v
	}

	fmt.Println(sum)

	c <- sum
}

// Channels are a typed conduit through which
// you can send and receive values with
// the channel operator, <-.
//
//  ch <- v    // Send v to channel ch.
//  v := <-ch  // Receive from ch, and
//             // assign value to v.
//
// (The data flows in the direction of the arrow.)
//
// Like maps and slices,
// channels must be created before use:
//
//  ch := make(chan int)
//
// By default, sends and receives block until
// the other side is ready. This allows goroutines
// to synchronize without explicit locks
// or condition variables.
//
// The example code sums the numbers in a slice,
// distributing the work between two goroutines.
// Once both goroutines have completed their
// computation, it calculates the final result.
func Channels() {
	n := []int{4, 3, -2, 0, 5, -8, -1, 1}

	c := make(chan int)
	go sum(n[:len(n)/2], c)
	go sum(n[len(n)/2:], c)
	go sum(n[6:], c)

	x, y, z := <-c, <-c, <-c

	fmt.Println(x, y, z)
}

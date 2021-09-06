package concurrency

import "fmt"

// https://tour.golang.org/concurrency/3
// Buffered Channels

// Channels can be buffered. Provide
// the buffer length as the second argument to
// make to initialize a buffered channel:
//
//  ch := make(chan int, 100)
//
// Sends to a buffered channel block only
// when the buffer is full. Receives block
// when the buffer is empty.
//
// Modify the example to overfill the buffer
// and see what happens.
func BufferedChannels() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	// ch <- 3 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

package concurrency

import "fmt"

// https://tour.golang.org/concurrency/5
// Select

func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// The select statement lets a goroutine wait
// on multiple communication operations.
//
// A select blocks until one of its cases can run,
// then it executes that case.
// It chooses one at random if multiple are ready.
func Select() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
}

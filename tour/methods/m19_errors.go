package methods

import (
	"fmt"
	"math"
	"time"
)

// https://tour.golang.org/methods/19
// Errors

type myError struct {
	time time.Time
	desc string
}

func (e *myError) Error() string {
	return fmt.Sprintf("at %v, %s", e.time, e.desc)
}

func run() error {
	return &myError{
		time.Now(),
		"opps!",
	}
}

// Go programs express error state
// with error values.
//
// The error type is a built-in interface
// similar to fmt.Stringer:
//
//  type error interface {
//      Error() string
//  }
//
// (As with fmt.Stringer,
// the fmt package looks for the error
// interface when printing values.)
//
// Functions often return an error value,
// and calling code should handle errors
// by testing whether the error equals nil.
//
//  i, err := strconv.Atoi("42")
//  if err != nil {
//      fmt.Printf("couldn't convert number: %v\n", err)
//      return
//   }
//  fmt.Println("Converted integer:", i)
//
// A nil error denotes success;
// a non-nil error denotes failure.
func Errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	math.Abs(100)
}

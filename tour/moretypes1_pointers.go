package tour

import "fmt"

// https://tour.golang.org/moretypes/1
// Pointers

// Go has pointers. A pointer holds the
// memory address of a value.
//
// The type *T is a pointer to a T value.
// Its zero value is nil.
//
//  var p *int
//
// The & operator generates a pointer
// to its operand.
//
//  i := 42
//  p = &i
//
// The * operator denotes the pointer's
// underlying value.
//
//  fmt.Println(*p) // read i through the pointer p
//  *p = 21         // set i through the pointer p
//
// This is known as "dereferencing"
// or "indirecting".
//
// Unlike C, Go has no pointer arithmetic.
func Pointers() {
	i, j := 4, 20

	p := &i                  // point to i
	fmt.Println("*p ->", *p) // read i through the pointer
	*p = 21                  // set i through the pointer
	fmt.Println("i ->", i)   // see the new value of i

	p = &j                 // point to j
	*p = *p / 10           // divide j through the pointer
	fmt.Println("j ->", j) // see the new value of j
}

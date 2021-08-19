package tour

import "fmt"

// https://tour.golang.org/moretypes/9
// Slice literals

// A slice literal is like an array
// literal without the length.
//
// This is an array literal:
//
//  [3]bool{true, true, false}
//
// And this creates the same array
// as above, then builds a slice that
// references it:
//
//  []bool{true, true, false}
//
func SliceLiterals() {
	numbers := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(numbers)

	primes := []bool{true, true, false, true, false, true}
	fmt.Println(primes)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
	}
	fmt.Println(s)

	st := []struct {
		i int
		b bool
	}{}

	for i := 0; i < len(numbers); i++ {
		st = append(st, struct {
			i int
			b bool
		}{numbers[i], primes[i]})
	}

	fmt.Println(st, " <- created with \"for\"")
}

package tourexercise

import (
	"golang.org/x/tour/reader"
)

// https://tour.golang.org/methods/22
// Exercise: Readers

type myReader struct{}

func (r myReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

// Solution of "readers" exercise
//
// For detail:
// https://tour.golang.org/methods/22
func Reader() {
	reader.Validate(myReader{})
}

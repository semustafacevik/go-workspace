package tourexercise

import (
	"io"
	"os"
	"strings"
)

// https://tour.golang.org/methods/23
// Exercise: rot13Reader

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	if err != nil {
		return n, err
	}

	for i := range b {
		if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
			b[i] += 13
		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
			b[i] -= 13
		} else {
			continue
		}
	}

	return n, err
}

// Solution of "rot13Reader" exercise
//
// For detail:
// https://en.wikipedia.org/wiki/ROT13
// https://tour.golang.org/methods/23
func RotReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

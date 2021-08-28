package tourexercise

import "golang.org/x/tour/pic"

// https://tour.golang.org/moretypes/18
// Exercise: Slices

func picture(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)

	for y := range p {
		p[y] = make([]uint8, dx)

		// try another combination
		// to see other images
		for x := range p[y] {
			p[y][x] = uint8(x)
			// p[y][x] = uint8(y)
			// p[y][x] = uint8(x + y)
			// p[y][x] = uint8(y - x)

			// p[y][x] = uint8((x + y) / 2)
			// p[y][x] = uint8(x * y)
			// p[y][x] = uint8(x ^ y)

			// p[y][x] = uint8((x ^ y) * (x ^ y))
			// p[y][x] = uint8(x*x + y*y)
		}
	}

	return p
}

// Solution of "slices" exercise
//
// For detail:
// https://tour.golang.org/moretypes/18
//
// For visual result:
// https://play.golang.org/p/14yIpsPP65v
func Slices() {
	pic.Show(picture)
}

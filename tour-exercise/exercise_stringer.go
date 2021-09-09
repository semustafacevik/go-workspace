package tourexercise

import "fmt"

// https://tour.golang.org/methods/18
// Exercise: Stringers

type ipAddress [4]byte

func (ip ipAddress) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// Solution of "stringers" exercise
//
// For detail:
// https://tour.golang.org/methods/18
func Stringer() {
	hosts := map[string]ipAddress{
		"loopback ": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

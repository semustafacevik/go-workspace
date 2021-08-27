package tourexercise

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

// https://tour.golang.org/moretypes/23
// Exercise: Maps

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		wc[word]++
	}

	fmt.Printf("%s\n%v\n\n", s, wc)

	return wc
}

// Solution of "maps" exercise
//
// For detail:
// https://tour.golang.org/moretypes/23
func Maps() {
	wc.Test(WordCount)
}

package methods

import "fmt"

// https://tour.golang.org/methods/17
// Stringers

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

// One of the most ubiquitous interfaces
// is Stringer defined by the fmt package.
//
//  type Stringer interface {
//      String() string
//  }
//
// A Stringer is a type that can
// describe itself as a string.
// The fmt package (and many others) look
// for this interface to print values.
func Stringer() {
	bs := person{"Busenaz Surmeneli", 23}
	mg := person{"Mete Gazoz", 22}

	fmt.Println(bs)
	fmt.Println(mg)
}

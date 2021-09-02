package methods

// https://tour.golang.org/methods/14
// The empty interface

// The interface type that specifies
// zero methods is known as the empty interface:
//
//  interface{}
//
// An empty interface may hold values
// of any type.
// (Every type implements at least zero methods.)
//
// Empty interfaces are used by code
// that handles values of unknown type.
// For example, fmt.Print takes
// any number of arguments of type interface{}.
func EmptyInterface() {
	var i interface{}
	describe(i)

	i = 8.21
	describe(i)

	i = "go-workspace"
	describe(i)
}
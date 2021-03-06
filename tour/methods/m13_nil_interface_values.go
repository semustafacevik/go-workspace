package methods

// https://tour.golang.org/methods/13
// Nil interface values

// A nil interface value holds
// neither value nor concrete type.
//
// Calling a method on a nil interface
// is a run-time error because there is
// no type inside the interface tuple to
// indicate which concrete method to call.
func NilInterfaceValues() {
	var i myInterface
	describe(i)
	// i.M() // run-time error
}

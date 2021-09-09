package basics

import "fmt"

// https://tour.golang.org/basics/10
// Short variable declarations

// Inside a function, the := short assignment
// statement can be used in place of a var
// declaration with implicit type.
//
// Outside a function, every statementbegins
// with a keyword (var, func, and so on) and so
// the := construct is n
func ShortVariableDeclarations() {
	var i, j int = -1, 0
	k := 4
	c, python, java, golang := true, false, 1, "studying"

	fmt.Println(i, j, k, c, python, java, golang)

}

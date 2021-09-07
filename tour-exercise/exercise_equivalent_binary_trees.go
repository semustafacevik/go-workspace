package tourexercise

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// https://tour.golang.org/concurrency/7
// Exercise: Equivalent Binary Trees

func walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkRecursive(t.Left, ch)
	ch <- t.Value
	walkRecursive(t.Right, ch)
}

func same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go walk(t1, ch1)
	go walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}

	return true
}

// Solution of "equivalent binary trees" exercise
//
// For detail:
// https://tour.golang.org/concurrency/7
func EquivalentBinaryTrees() {
	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println(same(tree.New(1), tree.New(2)))
}

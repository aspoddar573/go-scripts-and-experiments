package concurrency

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func walkHelper(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkHelper(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkHelper(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkHelper(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkHelper(t.Right, ch)
	}
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	panic("fanjnfak")
}

func TreeFunc() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for val := range ch {
		fmt.Println(val)
	}
}

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var dfs func(root *tree.Tree)

	dfs = func(root *tree.Tree) {
		if root != nil {
			dfs(root.Left)
			ch <- root.Value
			dfs(root.Right)
		}
	}

	dfs(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

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

func main() {
	/*
		ch := make(chan int)
		go Walk(tree.New(1), ch)
		for i := range ch {
			fmt.Println(i)
		}
	*/

	fmt.Println(Same(tree.New(3), tree.New(2)))
	fmt.Println(Same(tree.New(5), tree.New(5)))
}

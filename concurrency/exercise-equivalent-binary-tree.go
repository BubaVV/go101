package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var buf1, buf2 []int
	go func() {
		Walk(t1, ch1)
		defer close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		defer close(ch2)
	}()

	for x := range ch1 {
		buf1 = append(buf1, x)
	}

	for x := range ch2 {
		buf2 = append(buf2, x)
	}

	for i, v := range buf1 {
		if v != buf2[i] {
			return false
		}
	}
	return true
}

func main() {
	t := tree.New(2)
	ch := make(chan int, 10)
	fmt.Println(t)
	go func() {
		Walk(t, ch)
		defer close(ch)
	}()
	for x := range ch {
		fmt.Println(x)
	}
	fmt.Println("Should be true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should be false:", Same(tree.New(1), tree.New(2)))
}

// 1. Implement the Walk function.

// 2. Test the Walk function.

// The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

// Create a new channel ch and kick off the walker:

// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

// 4. Test the Same function.

// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

// The documentation for Tree can be found here.

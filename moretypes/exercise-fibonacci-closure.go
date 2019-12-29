package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	prelast := 0
	return func() int {
		if last == 0 {
			last = 1
			prelast = 0
		} else {
			prelast, last = last, last+prelast
		}
		return prelast

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

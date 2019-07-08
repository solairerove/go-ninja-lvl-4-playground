package jedi

import (
	"fmt"
)

// Ex1 fmt
func Ex1() {
	fmt.Println("Ex1")

	var x [5]int
	x[0] = 0
	x[1] = 1
	x[2] = 2
	x[3] = 3
	x[4] = 4

	fmt.Println(x)

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}

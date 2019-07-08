package jedi

import (
	"fmt"
)

// Ex2 fmt
func Ex2() {
	fmt.Println("Ex2")

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", x)
}

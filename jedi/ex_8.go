package jedi

import (
	"fmt"
)

// Ex8 fmt
func Ex8() {
	fmt.Println("Ex8")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}

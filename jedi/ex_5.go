package jedi

import (
	"fmt"
)

// Ex5 fmt
func Ex5() {
	fmt.Println("Ex5")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}

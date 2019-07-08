package jedi

import (
	"fmt"
)

// Ex7 fmt
func Ex7() {
	fmt.Println("Ex7")

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hello there"}

	xy := [][]string{x, y}

	fmt.Println(xy)

	for _, i := range xy {
		for _, k := range i {
			fmt.Println(k)
		}
	}
}

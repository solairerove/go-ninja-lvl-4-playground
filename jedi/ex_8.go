package jedi

import (
	"fmt"
)

// Ex8 fmt
func Ex8() {
	fmt.Println("Ex8")

	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range x {
		fmt.Printf("key: %v,\t value: %v\n", k, v)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

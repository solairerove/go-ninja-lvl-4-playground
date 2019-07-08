package jedi

import (
	"fmt"
)

// Ex9 fmt
func Ex9() {
	fmt.Println("Ex9")

	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["mikita"] = []string{"Kvas", "Pivo"}

	for k, v := range x {
		fmt.Printf("key: %v,\t value: %v\n", k, v)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

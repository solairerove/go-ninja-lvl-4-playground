package jedi

import (
	"fmt"
)

// Ex10 fmt
func Ex10() {
	fmt.Println("Ex10")

	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["mikita"] = []string{"Kvas", "Pivo"}

	delete(x, "mikita")

	for k, v := range x {
		fmt.Printf("key: %v,\t value: %v\n", k, v)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

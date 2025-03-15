package parser

import "fmt"

var depth int

var Active = false

func trace(name string) string {
	if Active {
		for i := 0; i < depth; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("BEGIN %s\n", name)
	}
	depth++
	return name
}

func untrace(name string) {
	if Active {
		for i := 0; i < depth-1; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("END %s\n", name)
	}
	depth--
}

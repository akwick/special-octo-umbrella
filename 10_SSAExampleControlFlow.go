package main

import "fmt"

func main() {
	i := 1
	s := "Hello I am an SSA gopher"
	if i < 1 {
		s = source()
	} else {
		s = "I am another SSA gopher"
	}
	sink(s)
}

// copied from 10_FuncAndVarDecl.go
func sink(s string) {
	fmt.Printf("A gopher reaches a sink: %s \n", s)
}

func source() string {
	return "I am an evil gopher"
}

package main

import "fmt"

func main() {
	i := 1
	s := source()
	if i < 1 {
		s = "I am the zeroth SSA gopher"
	} else {
		s = "I am not the zeroth SSA gopher"
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

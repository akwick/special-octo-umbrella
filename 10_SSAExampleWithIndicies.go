package main

import "fmt"

func main() {
	s1 := source()
	s2 := "I am an SSA gopher"
	sink(s2)
}

// copied from 10_FuncAndVarDecl.go
func sink(s string) {
	fmt.Printf("A gopher reaches a sink: %s \n", s)
}

func source() string {
	return "I am an evil gopher"
}

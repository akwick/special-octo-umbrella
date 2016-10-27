package main

import "fmt"

func main() {
	s := source()
	s = "I am a SSA gopher"
	sink(s)
}


// copied from 10_FuncAndVarDecl.go 
func sink(s string) {
	fmt.Printf("A gopher reaches a sink: %s \n", s)
}

func source() string{
	return "I am an evil gopher"
}

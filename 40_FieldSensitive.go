package main

import "fmt"

type SecretGopher struct {
	private, public string
}

func main() {
	s1 := SecretGopher{source(), "XXX"}
	s2 := SecretGopher{"XXX", "XXX"}
	sink(s1.public)
	sink(s2.private)
}

// copied from 10_FuncAndVarDecl.go
func sink(s string) {
	fmt.Printf("A gopher reaches a sink: %s \n", s)
}

func source() string {
	return "I am an evil gopher"
}

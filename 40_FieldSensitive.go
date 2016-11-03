package main

import "fmt"

type SecretGopher struct {
	secret, public string
}

func main() {
	s1 := SecretGopher{source(), source()}
	s2 := SecretGopher{source(), "XXX"}
	sink(s1.public)
	sink(s2.public)
}

// copied from 10_FuncAndVarDecl.go
func sink(s string) {
	fmt.Printf("A gopher reaches a sink: %s \n", s)
}

func source() string {
	return "I am an evil gopher"
}

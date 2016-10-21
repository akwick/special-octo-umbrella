package main

import "fmt"

func main() {
	var iptr, jptr *int
	i := 42
	iptr = &i
	jptr = new(int)
	fmt.Printf("i:     %d\n", i)
	fmt.Printf("iptr:  %d\n", iptr)
	fmt.Printf("*iptr: %d\n", *iptr)
	fmt.Printf("jptr:  %d\n", jptr)
	fmt.Printf("*jtpr: %d\n", *jptr)
}

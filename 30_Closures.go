package main

import "fmt"

func main() {
	g := &gopher{"green"}
	sayHelloToClosures := func() {
		g.sayHello()
	}
	sayHelloToClosures()
}

// Repeating definition from previous examples to cut the file here.
// No guarantee that it's always the same as in the previous examples.
type gopher struct {
	color string
}

func (g gopher) sayHello() {
	fmt.Printf("Hello I am a %s gopher\n", g.color)
}

package main

import "fmt"

func CreateEvilGopherConcurrent(c chan EvilGopher) {
	g := EvilGopher{evilHabit: "not care about synchronization", gopher: gopher{"purple"}}
	c <- g
}

func main() {
	c := make(chan EvilGopher)
	go CreateEvilGopherConcurrent(c) 
	g :=<- c
	g.DoEvilStuff()
}


// Repeating definition from previous examples to cut the file here.
// No guarantee that it's always the same as in the previous examples.
type gopher struct {
	color string
}

func (g gopher) sayHello() {
	fmt.Printf("Hello I am a %s gopher\n", g.color)
}

type EvilGopher struct {
	gopher
	evilHabit string
}

func (e EvilGopher) DoEvilStuff() {
	e.sayHello()
	fmt.Printf("The %s evil gopher does: %s\n", e.gopher.color, e.evilHabit)
}

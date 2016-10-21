package main

import "fmt"

func (e *EvilGopher) DoAnotherEvilStuff(s *string) {
	e.evilHabit = *s
	*s = *s + " and changes values"
	e.DoEvilStuff()
}

func main() {
	g := EvilGopher{evilHabit: "not using pointer receivers", gopher: gopher{"purple"}}
	g.DoEvilStuff()
	s := "not knowing what pointer receivers do"
	g.DoAnotherEvilStuff(&s)
	fmt.Printf("s: %s\n", s)
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

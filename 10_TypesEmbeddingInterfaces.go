package main

import "fmt"

type gopher struct {
	color string
}

func (g gopher) sayHello() {
	fmt.Printf("Hello I am a %s Gopher\n", g.color)
}

type evilGopher struct {
	gopher
	evilHabit string
}

func (e evilGopher) doEvilStuff() {
	e.sayHello()
	fmt.Printf("The %s evil Gopher does: %s\n", e.gopher.color, e.evilHabit)
}

type lovelyGopher struct {
	g gopher
}

func (l lovelyGopher) sayHello() {
	fmt.Printf("Hello I am a lovely %s Gopher\n", l.g.color)
}

type helloer interface {
	sayHello()
}

func main() {
	g1 := gopher{color: "blue"}
	g2 := evilGopher{evilHabit: "not using gofmt", gopher: gopher{"pink"}}
	g3 := lovelyGopher{gopher{"green"}}

	var hellos = []helloer{g1, g2, g3}
	for _, h := range hellos {
		h.sayHello()
	}
	g2.doEvilStuff()
}

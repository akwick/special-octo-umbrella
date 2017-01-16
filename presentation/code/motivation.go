package main

func main() {
	a := "Hello Gopher"
	ch := make(chan string)
	go f(ch)
	sink(&x)
	a = source()
	ch <- x
}

func f(ch chan string) {
	b := <- ch
	sink(&y)
}

func sink(s *string) {}

func source() string {
	return "secret"
}

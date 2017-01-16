package main

import "fmt"

func main() {
	msgs := make(chan string, 0)
	fmt.Println("Hello")
	go hello(msgs)
	msg := <-msgs
	fmt.Println(msg)
}

func hello(c chan string) {
	c <- "Hello Gopher"
}

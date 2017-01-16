package main

import (
	"fmt"
	"time"
)

func main() {
	msgs := make(chan string, 0)
	go func() { 
		time.Sleep(3000 * time.Millisecond)
		msgs <- "Hello Gopher" }()  
	msg := <-msgs
	fmt.Println(msg)
}

package main

import "fmt"

func main() {
	event := "Go meetup Frankfurt"
	success := true
	_ = "breakpoint"

	fmt.Printf("The type of event is: %T \n", event)
	fmt.Printf("The type of success is: %T \n", success)
}

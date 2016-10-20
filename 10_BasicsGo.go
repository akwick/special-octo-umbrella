package main
import "fmt"

func sink(s string) {
  fmt.Printf("A gopher reaches a sink: %s \n", s)
}
func source() string{
  return "I am an evil gopher"
}

func main() {
  var a string
  a = "I am a lovely gopher"
  sink(a)
  b := source()
  sink(b)
}

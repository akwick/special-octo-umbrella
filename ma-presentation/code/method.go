package main

import "fmt"

// START OMIT
type Password struct {p string}

func CheckPassword(a Password, b Password) bool{
  return a == b
} 

func  (p Password) CheckPassword(a Password) bool { 
  return p == a
}
// END OMIT

func main() {
  a := Password{"Hello"}
  b := Password{"Hello"}
  c := Password{"Gopher"}
  fmt.Printf("a match b %s  | %s\n", CheckPassword(a, b), a.CheckPassword(b))
  fmt.Printf("a match c %s  | %s\n", CheckPassword(a, c), a.CheckPassword(c))
}

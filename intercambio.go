package main

import "fmt"
func intercambio(a *int, b *int){
  aux_a := *a
  aux_b := *b

  *a = aux_b
  *b = aux_a

}

func main() {
  a := 7
  b := 2

  intercambio(&a, &b)

  fmt.Println(a)
  fmt.Println(b)

}
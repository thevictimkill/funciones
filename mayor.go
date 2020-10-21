package main

import "fmt"

func sum(args ...int) int {
	actual := -5
 	for _, v := range args { 
		if actual < v {
      actual = v
    }
	}

	return actual
}

func main() {
  a := []int{1, 4, 2, 99, 9, 5, 7}
  max := sum(a...)
  fmt.Println(max)
}
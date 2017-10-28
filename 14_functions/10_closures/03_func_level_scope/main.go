package main

import "fmt"

var x int

func main() {

	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

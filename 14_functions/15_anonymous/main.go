package main

import (
	"fmt"
)

func main() {
	func(x int){
		fmt.Println("Hello, Anonymous")
		fmt.Println(x)
	}(42)
}

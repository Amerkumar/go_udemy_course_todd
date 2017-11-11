package main

import "fmt"

func main() {
	//Only way to add a function inside a function
	greeting := func() { fmt.Println("Hello World!") }
	greeting()

}

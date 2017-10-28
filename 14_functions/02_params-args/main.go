package main

import "fmt"

func main() {
	greet("Jane") // argument
	greet("John") // parameter

}

func greet(name string){ // declared with parameter - recives an argument
	fmt.Println(name)
}
package main

import "fmt"

func main() {

	//no default fallthrough in go
	//fallthrough is optional

	switch "a" {
	case "A", "a":
		fmt.Println("Grade A")
	case "B", "b":
		fmt.Println("Grade B")
	case "C", "c":
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

}

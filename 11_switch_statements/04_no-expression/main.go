package main

import "fmt"

func main() {

	//no default fallthrough in go
	//fallthrough is optional
	name := "AmerK"

	switch {
	case len(name) == 4:
		fmt.Println("Grade A")
	case name == "AmerK":
		fmt.Println("Grade B")

	default:
		fmt.Println("F")
	}

}

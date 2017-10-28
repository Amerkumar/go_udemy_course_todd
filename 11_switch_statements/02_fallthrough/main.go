package main

import "fmt"

func main() {

	//no default fallthrough in go
	//fallthrough is optional


	switch "A" {
	case "A":
		fmt.Println("Grade A")
		fallthrough
	case "B":
		fmt.Println("Grade B")
	case "C":
		fmt.Println("C")
	default:
		fmt.Println("F")
	}

}


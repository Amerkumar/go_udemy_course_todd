package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {

	switch x.(type) { //This is an assertion :
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	case int32:
		fmt.Println("rune")
	default:
		fmt.Println("unknown")
	}

}
func main() {
	SwitchOnType(7)
	SwitchOnType("Amer")
	var t = Contact{"Hello", "Ahmed"}
	SwitchOnType(t)
	SwitchOnType('a')

}

package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is happy"
		fmt.Println(y)
	}
	// fmt.Println(y) // cannot do it here - outside scope of y
}

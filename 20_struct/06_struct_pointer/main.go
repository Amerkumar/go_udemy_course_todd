package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := &person{"James", "Bond"}
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
}

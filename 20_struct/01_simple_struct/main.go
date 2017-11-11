package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 80}
	var p3 person
	var p4 person
	p3.first = "Amer"
	p3.last = "Kumar"
	p3.age = 22
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
}

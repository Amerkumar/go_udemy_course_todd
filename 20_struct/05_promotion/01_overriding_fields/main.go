package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	person
	first         string
	LicenseToKill bool
}

func main() {

	p1 := DoubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   20,
		},
		first:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		person: person{
			first: "Money",
			last:  "Penny",
			age:   19,
		},
		first:         "If looks could kill",
		LicenseToKill: false,
	}
	//Outer Type fields overwrite inner type fields
	fmt.Println(p1.first, p1.person.first)
	fmt.Println(p2.first, p2.person.first)
}

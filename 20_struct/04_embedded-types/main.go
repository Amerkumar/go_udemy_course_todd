package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type DoubleZero struct {
	person
	LicenseToKill bool
}

func main() {

	p1 := DoubleZero{
		person: person{
			first: "James",
			last:  "Bond",
			age:   20,
		},
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		person: person{
			first: "Money",
			last:  "Penny",
			age:   19,
		},
		LicenseToKill: false,
	}

	p3 := DoubleZero{person{"Amer", "kumar", 22}, true}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}

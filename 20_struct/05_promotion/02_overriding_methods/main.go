package main

import "fmt"

type person struct {
	first string
	last  string
}

type DoubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss MoneyPenny, so good to see you,")
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := DoubleZero{
		person: person{
			first: "Money",
			last:  "Penny",
		},
		LicenseToKill: false,
	}
	//Outer Type fields overwrite inner type fields
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}

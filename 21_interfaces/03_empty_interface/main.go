package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{
		animal : animal{
			sound : "meow",
			},
			annoying : false,
	}

	critters := []interface{}{fido,fifi}
	fmt.Println(critters)

}

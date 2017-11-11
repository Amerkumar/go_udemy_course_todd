package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first    string
	last     string
	age      int
	Exported int
}

func main() {
	p1 := person{"James", "Bond", 22, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

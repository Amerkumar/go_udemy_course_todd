package main

import "fmt"

func main() {

	age := 44
	changeMe(age)
	fmt.Println(age) //24

}

func changeMe(z int)  {
	fmt.Println(z) // 0xc04200a290
	z = 24
}

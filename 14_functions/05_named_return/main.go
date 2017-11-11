package main

import "fmt"

func main() {
	fmt.Println(greet("Amer ", "Kumar"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

package main

import "fmt"

func makeIterator(s []string) func() func() string {
	i := 0
	return func() func() string {
		if i == len(s) {
			return nil
		}
		j := i
		i++
		return func() string {
			return s[j]
		}
	}
}

func main() {

	i := makeIterator([]string{"hello", "world", "this", "is", "dog"})

	fmt.Printf("%T \n",i)
	for c := i(); c != nil; c = i() {
		fmt.Printf("%T -> ",c)
		fmt.Println(c())
	}

}
package main

import "fmt"

func main() {

	mySlice := []int{1,3,5,7,9,11} // Slices do not have number between brackets
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}

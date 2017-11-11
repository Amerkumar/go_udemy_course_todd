package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}

	mySlice = append(mySlice[0:1], mySlice[2:]...)

	fmt.Println(mySlice)
}

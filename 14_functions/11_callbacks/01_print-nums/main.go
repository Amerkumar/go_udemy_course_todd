package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})
}

/*
1. Goes into main
2. Gives callback definition of func - Here
callback =  func(n int) {
		fmt.Println(n)
	}
func visit(numbers []int, callback func(int))  {
	for _, n := range numbers{
		callback(n)
	}
}

*/

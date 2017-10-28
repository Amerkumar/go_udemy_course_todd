package main

import "fmt"

func main() {
	//break use
	i = 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

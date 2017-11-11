package main

import "fmt"

func main() {

	//nested loop
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		} //end of inner loop
	} //end of outer loop

}

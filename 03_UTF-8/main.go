package main

import "fmt"

func main()  {
	//%q for UTF - 8 chracters
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i)
	}

}
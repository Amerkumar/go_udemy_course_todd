package main

import "fmt"

//Three threads func main thread exits promptly

func main()  {
	go foo()
	go bar()
}

func foo()  {
	for i := 0; i < 100; i++{
		fmt.Println("Foo: ",i)
	}
}

func bar()  {
	for i := 0; i < 100; i++{
		fmt.Println("Bar:", i)
	}
}

package main

import (
	"sync"
	"fmt"
	"time"
	"runtime"
)

var wg sync.WaitGroup

//Init is the function that run before main and only for first time.
func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo()  {
	for i := 0; i < 100; i++{
		fmt.Println("Foo: ",i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar()  {
	for i := 0; i < 100; i++{
		fmt.Println("Bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

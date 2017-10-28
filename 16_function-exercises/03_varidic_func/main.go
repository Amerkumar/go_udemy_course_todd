package main

import "fmt"

func greatest(a ...int) int {
	max := a[0]
	for i := 1; i < len(a); i++{
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func main() {
		fmt.Println(greatest(1,10,3,4))

		//This works also
		b := []int{1,2,3,4,13,123}
		fmt.Println(greatest(b...))
		//But this does not work
		//fmt.Println(greatest(b))
	}

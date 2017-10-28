package main

import "fmt"

func main() {
	n := average(43, 57, 87, 12, 45, 57)
	fmt.Println(n)
}

//Calculates the average of the list of numbers
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	for _, v := range sf { // range gives index , value
		total += v
	}
	return total / float64(len(sf))
}

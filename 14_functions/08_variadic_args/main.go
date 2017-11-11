package main

import "fmt"

func main() {

	//These two are different types
	data := []float64{43, 57, 87, 12, 45, 57} //Slice of different types
	n := average(data...)                     //We are giving variadic argument form
	fmt.Println(n)
}

//Calculates the average of the list of numbers
func average(sf ...float64) float64 { // sf []float64 in case of data
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	for _, v := range sf { // range gives index , value
		total += v
	}
	return total / float64(len(sf))
}

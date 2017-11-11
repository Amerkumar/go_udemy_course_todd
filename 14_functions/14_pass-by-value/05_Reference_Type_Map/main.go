package main

import "fmt"

func main() {
	z := make(map[string]int)
	changeMe(z)
	fmt.Println(z["Todd"])
}

func changeMe(x map[string]int) {
	x["Todd"] = len("Todd")
}

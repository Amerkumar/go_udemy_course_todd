package main

import "fmt"

func main() {

	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"
	myGreeting["Harleen"] = "Howdy!!!"
	myGreeting["DumbleDore"] = "Gidday!"

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}

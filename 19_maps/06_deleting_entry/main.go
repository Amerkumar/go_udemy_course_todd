package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"

	myGreeting["Harleen"] = "Howdy!!!"

	fmt.Println(len(myGreeting))

	delete(myGreeting,"Tim")

	fmt.Println(myGreeting)


}

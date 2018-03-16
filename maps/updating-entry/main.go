package main

import "fmt"

func main() {
	myGreeting := map [string]string{
		"Tim": "Good Morning",
		"Jenny": "Bonjour",
		"Chris": "Hey-OH!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Tim"])
	fmt.Println(myGreeting["Jenny"])
	fmt.Println(myGreeting["Chris"])
	fmt.Println(myGreeting["Harleen"])

	myGreeting["Harleen"] = "Hola!"
	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Harleen"])
	}
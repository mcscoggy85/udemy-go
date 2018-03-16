package main

import "fmt"

func main() {
	myGreeting := map [string]string{
		"Tim": "Good Morning",
		"Jenny": "Bonjour",
		"Chris": "Hey-OH!",
	}

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Tim"])
	fmt.Println(myGreeting["Jenny"])
	fmt.Println(myGreeting["Chris"])
}
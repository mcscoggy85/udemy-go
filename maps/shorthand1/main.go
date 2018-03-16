package main

import "fmt"

func main() {
	myGreeting := make(map [string]string)
	//alternative way below
	//myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Tim"])
	fmt.Println(myGreeting["Jenny"])
}
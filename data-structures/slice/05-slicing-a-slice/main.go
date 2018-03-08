package main

import "fmt"

func main() {

	greeting := []string{
		"Good Morning!",
		"Bonjour!",
		"dias!",
		"Whats up",
		"How are you",
		"yo",
		"sup",
		"how you doin",
		"hey",
	}

	fmt.Println("[1:2]", greeting[1:2])
	fmt.Println("[:2]", greeting[:2])
	fmt.Println("[5:]", greeting[5:])
	fmt.Println("[:]", greeting[:])
}

package main

import (
	"fmt"
)

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning"
	greeting[1] = "Hello"
	greeting[2] = "Good night"
	greeting = append(greeting, "hola")
	greeting = append(greeting, "yello")
	greeting = append(greeting, "whats up")
	greeting = append(greeting, "tata")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
package main

import "fmt"

func main() {

	greeting := []string {
		"Good Morning!",
		"Bonjour!",
		"dias!",
	}

	for i, value := range greeting {
		fmt.Println(i, "==",value)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}

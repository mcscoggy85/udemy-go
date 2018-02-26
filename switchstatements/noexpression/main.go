package main

import . "fmt"

func main() {
	friendsname := "Chris"

	switch { //conditional statement
	case friendsname == "Daniel":
		Println("Wassup Daniel")
	case len(friendsname) == 3:
		Println("Wassup Medhi")
	case friendsname == "Jenny":
		Println("Wassup Jenny")
	default:
		Println("Have you no friends?...")
	}
}
package main

import . "fmt"

func main() {
	switch "Medhi"{ //conditional statement
	case "Daniel":
		Println("Wassup Daniel")
	case "Medhi":
		Println("Wassup Medhi")
	case "Jenny":
		Println("Wassup Jenny")
	default:
		Println("Have you no friends?...")

// all case statements are options if the condition is met
// if no conditions are met the default runs just like an else statement
	}
}

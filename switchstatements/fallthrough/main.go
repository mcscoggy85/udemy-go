package main

import . "fmt"

func main() {
	switch "Medhi" { //conditional statement
	case "Daniel":
		Println("Wassup Daniel")
	case "Medhi":
		Println("Wassup Medhi")
		fallthrough
	case "Jenny":
		Println("Wassup Jenny")
	case "Julian":
		Println("Wassup Julian")
	default:
		Println("Have you no friends?...")
	}
}
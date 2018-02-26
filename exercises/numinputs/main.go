package main

import . "fmt"

func main() {
	var num1 int
	var num2 int

	Println("I am going to ask you for two numbers...")
	Printf("What is the first number you want? Enter here: \n\n")
	Scan(&num1)
	Printf("What is the second number you want? Enter here: \n\n")
	Scan(&num2)

	Printf("The numbers you chose were %v, and %v\n", num1, num2)

	Println(num1, "%", num2, "=", num1%num2)

}
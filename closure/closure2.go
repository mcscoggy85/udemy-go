package main

import "fmt"

func main() {
	x := 0 //narrow scope instead of using a global scope variable
	increment := func() int{  // This line is a nested function in go as well as an anonymous function
		x++					  // an anon function is a function without a name, it looks like a function that is
		return x			  // assigned to a variable *** a func expression ***
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	}
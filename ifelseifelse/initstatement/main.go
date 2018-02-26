package main

import . "fmt"

func main() {
	b := true

	if food := "Chocolate"; b { //forces true with initialized expression of b
		Println(food)			// this practice narrows the scope of the variables
								// being used in other places
	}
}

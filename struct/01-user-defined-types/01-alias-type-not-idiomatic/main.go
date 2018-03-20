package main

import (
	"fmt"
)

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Println(myAge)
	fmt.Printf("%T %v \n", myAge, myAge)
}
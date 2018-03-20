package main

import "fmt"

type Person struct {
	first string
	last string
	age int
}

func main() {
	p1 := Person{"Chris", "Scogin", 29}
	p2 := Person{"Charlie", "Stamp", 34}
	fmt.Println(p1.first, p1.last, "is",p1.age)
	fmt.Println(p2.first, p2.last, "is", p2.age)
}
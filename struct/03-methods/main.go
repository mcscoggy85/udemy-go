package main

import (
	"fmt"
)

type person struct {   //example of type class
	first string
	last string
	age int
}

func (p person) fullName() string {  //example of a method
	return p.first + " " + p.last

}

func main() {
	p1 := person{"Chris", "Scogin", 29}
	p2 := person{"Charlie", "Stamp", 34}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p1.age)
}

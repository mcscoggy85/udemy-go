package main

import "fmt"

const (
	_ = iota
	Kb = 1 << (iota * 10)
	Mb = 1 << (iota * 10)
)

func main() {
	fmt.Println(Kb)
	fmt.Println(Mb)
	fmt.Println("Binary\t\t\t\t\tDecimal")
	fmt.Printf("%b\t\t\t\t", Kb)
	fmt.Printf("%d\n", Kb)
	fmt.Printf("%b\t", Mb)
	fmt.Printf("%d\n", Mb)

}
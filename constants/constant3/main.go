package main

import "fmt"

// iota is a small amount
const (
	A = iota // 0
	B
	C
)

const (
	D = iota // 0
	E
	F
)
const (
	_ = iota
	G = iota * 10
	H = iota * 10
)

func main(){
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)

	fmt.Println(G)
	fmt.Println(H)

}
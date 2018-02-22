package main

import (
	"fmt"
)

func main(){
	x := 13 % 3
	fmt.Println(x)
	fmt.Println()
	evenorodd()
}

func evenorodd() {
	for i := 1; i < 11; i++ {
		// fmt.Println(i)
		if i % 2 == 1{
			//fmt.Println("Odd")
			fmt.Printf("%d = Odd\n", i)
		} else {
			fmt.Printf("%d = Even\n", i)
			//fmt.Println("Even")
		}
	}
}
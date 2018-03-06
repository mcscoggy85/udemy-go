package main

import (
	"fmt"
)

func takesint(n int) {
	arg1 := n / 2
	fmt.Println(arg1)

	if n % 2 == 0{
		fmt.Println("Even")
	} else {
		fmt.Println("False")
	}

}

func half(n int) (int, bool) {
return n / 2, n % 2 == 0
}

func main()  {
	takesint(4)

	h,even := half(10)
	fmt.Println("(",h,")","(",even,")")
}
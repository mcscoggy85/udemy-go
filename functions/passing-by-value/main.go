package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age) //mem addr

	changeMe(&age)

	fmt.Println(&age) //mem addr
	fmt.Println(age) //value

}

func changeMe(z *int)  {
	fmt.Println(z) //mem addr
	fmt.Println(*z) //pointer to value
	*z = 24
	fmt.Println(z) //mem addr
	fmt.Println(*z) //pointer to value

}

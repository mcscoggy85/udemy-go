package main

import "fmt"

func hello() {
	fmt.Println("Hello ")
}

func world() {
	fmt.Println("World")
}

func main()  {
	defer world()  //will defer until right before main func ends
	hello()

}
package main

import "fmt"

func main() {
	myslice := []int {1,3,5,7,9,11}
	for i, value := range myslice {
		fmt.Println(i, "==", value)
	}
	}


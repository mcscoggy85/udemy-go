package main

import "fmt"

func greatnum(n ...int) int{
	var largest int
	for _,v := range n {
		if v > largest {
			largest = v
		}
	}
	return largest

}

func main() {
	greatest := greatnum(4, 7, 8, 3, 10, 11, 125, 4)
	fmt.Println(greatest)
}
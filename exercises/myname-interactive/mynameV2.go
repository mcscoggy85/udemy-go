package main

import . "fmt"

func main() {
	var myname string
	Printf("What is your name? ")
	Scan(&myname)
	Printf("Your name is %v", myname)
}
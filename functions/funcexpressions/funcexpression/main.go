package main

import "fmt"

func main() {
	greeting := func() { //func expression example also known as an anon function
		fmt.Println("Hello World") //func assigned to a variable
	}
	greeting()
}


package main

import "fmt"

func main()  {

	i := 1 //initializer

	for {
		fmt.Println(i)
		if i >= 10{  //condition
			break
		}
		i++  //post
		}
}
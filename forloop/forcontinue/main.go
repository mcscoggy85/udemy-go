package main

import "fmt"

func main()  {

	i := 0 //initializer

	for{
		i++ //no longer a post its now an incrementer this incrementer can also exist
		// as a post at the end with a slight change
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >=50 {
			break
		}
	}

}
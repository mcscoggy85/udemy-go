package main

import (
	. "fmt"
)

func main()  {
	if !true {  //synonym for false
		Println("This will not run")
	}

	if !false {  //synonym for true
		Println("This will run")
	}
}


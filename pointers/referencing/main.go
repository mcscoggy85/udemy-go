package main

import (
	"fmt"
	)

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	//var b *int = &a
	//var b = &a
	var b = a
	fmt.Println(b)

	// above code makes "b" a pointer to mem addr where an int is stored
	// b is of type "int" pointer
	// *int -- the * is part of the type -- "b" is of type *int
}
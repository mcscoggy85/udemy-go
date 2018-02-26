package main

import ( .
	"fmt"
)

type Contact struct {
	greeting string
	name string
	}

func SwitchOnType(x interface{})  {
	switch x.(type) { //This is an assert; asserting, "x is of this type"
	case int:
		Println("int")
	case string:
		Println("String")
	case Contact:
		Println("Contact")
	default:
		Println("Unknown")
	}

}

func main()  {
	SwitchOnType(7)
	SwitchOnType("Scogin")
	var t = Contact{"Good to see you,", "Tim"}
	SwitchOnType(t)

	}

package main

import "fmt"

const metersToYards = 1.09361

func main()  {
	var meters float64
	fmt.Print("Enter Meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.")
}
package main

import "fmt"

func main() {
	// Create list of day
	var days = [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	var slice24 = days[2:4]
	fmt.Println(slice24)      // [Tuesday Wednesday] Get data between index 2 - 4
	fmt.Println(len(slice24)) //2 // length data
	fmt.Println(cap(slice24)) //5 // capacity total from Tuesday until Saturday

	// Update data from array days will be effect to slice24
	// days[3] = "WednesdayX"
	// fmt.Println(slice24)

	// Update data from slice will be effect to array days to
	// Index 0 from slice
	// slice24[0] = "TuesdayX"
	// fmt.Println(days)

	// Create new slice
	var slice5 = days[5:]
	fmt.Println("slice5 : ", slice5)

	// Append if slice is full capacity
	slice5 = append(slice5, "New day")
	fmt.Println("slice5 : ", slice5) //[Saturday New day] Will create new array
	fmt.Println("days : ", days)     // array days not effected because when append slice is full capacity

	var slice46 = days[4:6]
	fmt.Println("slice46 : ", slice46)

	// Append if slice is not full capacity
	slice46 = append(slice46, "New Saturday") /// Will replace index 6
	fmt.Println("slice46 : ", slice46)
	fmt.Println("days : ", days) //[Sunday Monday TuesdayX Wednesday Thursday Friday New Saturday]
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Get time now
	now := time.Now()

	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Local())

	// Create date
	createDate := time.Date(1992, time.July, 10, 4, 30, 0, 0, time.Local)
	fmt.Println(createDate.Local())

	// Parse string to date
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1992-07-10")
	fmt.Println(parse)
}

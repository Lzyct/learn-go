package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBool, err := strconv.ParseBool("1")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parseBool) // true
	}

	// using decimal format
	parseInt, err := strconv.Atoi("60000")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(parseInt)
	}

	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	// using decimal format
	formatInt := strconv.Itoa(696969)
	fmt.Println(formatInt)
}

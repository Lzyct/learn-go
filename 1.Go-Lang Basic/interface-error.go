package main

import (
	"errors"
	"fmt"
)

func divide(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("Divider can't 0")
	} else {
		return value / divider, nil
	}
}
func main() {
	result, error := divide(10, 1)

	if error == nil {
		fmt.Println(result)
	} else {
		fmt.Println(error)
	}
}

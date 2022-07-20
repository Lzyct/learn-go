package main

import "fmt"

func randomType() interface{} {
	return 2
}
func main() {
	random := randomType()

	// Make sure the type is same
	// var style1 string = random.(string)
	// fmt.Println(style1)

	switch style2 := random.(type) {
	case string:
		fmt.Println("Value", style2, "is string")
	case int:
		fmt.Println("Value", style2, "is int")
	default:
		fmt.Println("Value unknown type")
	}
}

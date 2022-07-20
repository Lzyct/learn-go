package main

import "fmt"

func main() {
	count := 1

	// 1
	for count <= 10 {
		fmt.Println("Count Method 1 : ", count)
		count++
	}

	// 2
	for i := 1; i <= 10; i++ {
		fmt.Println("Count Method 2 : ", i)
	}

	// Loop from slice data
	slice := []string{"a", "b", "c", "d"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("1 -> index : ", i, ",value : ", slice[i])
	}

	for i, value := range slice {
		fmt.Println("2 -> index : ", i, ",value : ", value)
	}

	// Loop from map data
	student := make(map[string]string)
	student["name"] = "Lzyct"
	student["rank"] = "1"
	student["class"] = "SS"

	for key, value := range student {
		fmt.Println("Key :", key, ", value : ", value)
	}

	for _, value := range student {
		fmt.Println("Value : ", value)
	}

}

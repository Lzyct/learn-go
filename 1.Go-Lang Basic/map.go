package main

import "fmt"

func main() {
	// Declare 1
	map1 := map[string]string{
		"title":       "Learn Go",
		"description": "with Lazycat Labs",
	}
	fmt.Println(map1)
	fmt.Println(map1["title"])
	fmt.Println(map1["description"])

	// Declare 2
	map2 := make(map[string]string)
	map2["title"] = "Learn Go"
	map2["description"] = "with Lazycat Labs"

	fmt.Println(map2)
	fmt.Println(map2["title"])
	fmt.Println(map2["description"])
	fmt.Println(len(map2))

	// Delete data on map2
	delete(map2, "title")
	fmt.Println(map2)
	fmt.Println(len(map2))
}

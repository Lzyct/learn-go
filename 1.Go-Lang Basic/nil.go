package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}
}

func main() {
	newMap := newMap("Lzyct")

	if newMap == nil {
		fmt.Println("No Data")
	} else {
		fmt.Println(newMap)
	}

}

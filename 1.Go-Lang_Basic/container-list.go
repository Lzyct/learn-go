package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkedList := list.New()

	linkedList.PushBack("1")
	linkedList.PushBack("2")
	linkedList.PushBack("3")

	// Push data to the first
	linkedList.PushFront("0")
	linkedList.PushFront("-1")

	fmt.Println("Get the first data : ", linkedList.Front().Value)
	fmt.Println("Get the latest data : ", linkedList.Back().Value)

	// Get all data from the first
	for element := linkedList.Front(); element != nil; element = element.Next() {
		fmt.Println("element >-> :", element.Value)
	}

	// Get all data from the last
	for element := linkedList.Back(); element != nil; element = element.Prev() {
		fmt.Println("element <-< :", element.Value)
	}
}

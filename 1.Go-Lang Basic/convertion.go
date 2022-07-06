package main

import "fmt"

func main() {
	name := "Lzyct"

	fmt.Println(name)
	fmt.Println("Lenght from Lzyct ", len(name)) //5
	fmt.Println("Index 0 from Lzyct", name[0])   // 76 : byte number for L
	fmt.Println("Index 1 from Lzyct", name[1])   // 122 : byte number for z

	fmt.Println("L from Lzyct", string(name[0])) // 76 : byte number for L
	fmt.Println("z from Lzyct", string(name[1]))
}

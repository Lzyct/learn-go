package main

import "fmt"

type GithubUser struct {
	UserName, Name string
	Stars          int
}

func main() {
	var style1 GithubUser
	style1.UserName = "Lzyct"
	style1.Name = "Mudassir"
	style1.Stars = 69
	fmt.Println("Style 1 :", style1)

	style2 := GithubUser{
		UserName: "Lzyct",
		Name:     "Mudassir",
		Stars:    69,
	}

	fmt.Println("Style 2 :", style2)

	style3 := GithubUser{"Lzyct", "Mudassir", 69}
	fmt.Println("Style 3 :", style3)
}

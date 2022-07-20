package main

import "fmt"

type GithubUser struct {
	UserName, Name string
	Stars          int
}

/**
Struct function, is like method inside struct
but actually it's same like this

func structFunction(githubUser GithubUser,name string)

structFunction(style2,"ukieTux")

You can using struct function to create function for specific struct.
*/
func (githubUser GithubUser) structFunction(name string) {
	fmt.Println("Hello", name, "Please follow", githubUser.UserName, "on Github")
}

func main() {
	var style1 GithubUser
	style1.UserName = "Lzyct"
	style1.Name = "Mudassir"
	style1.Stars = 69
	// fmt.Println("Style 1 :", style1)

	style2 := GithubUser{
		UserName: "Lzyct",
		Name:     "Mudassir",
		Stars:    69,
	}

	fmt.Println("Style 2 :", style2)
	style2.structFunction("ukieTux")

	// style3 := GithubUser{"Lzyct", "Mudassir", 69}
	// fmt.Println("Style 3 :", style3)
}

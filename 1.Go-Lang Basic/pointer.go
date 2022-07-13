package main

import "fmt"

type GithubData struct {
	UserName, Name, Email string
}

/**
Recommended using pointer struct on function
to avoid memory leak
*/
func updateEmail(githubData *GithubData) {
	githubData.Email = "ukie.tux@gmail.com"
}

func main() {
	// by default in Go Lang is "pass by value"
	user1 := GithubData{"Lzyct", "Mudassir", "hey.mudassir@gmail.com"}
	// user2 duplicate value from user1
	// user2 := user1

	//if you change user2, the data on user1 not changed
	// user2.UserName = "Lazycat"

	// fmt.Println("user1 :", user1) // {Lzyct Mudassir hey.mudassir@gmail.com}
	// fmt.Println("user2 :", user2) // {Lazycat Mudassir hey.mudassir@gmail.com}

	// we can using "pass by reference" using pointer
	// now, user3 and user4 reference to user1
	user3 := &user1
	user4 := &user1

	fmt.Println("user1 :", user1)
	fmt.Println("user4 :", user4)

	fmt.Println("===============================")

	/**
	if you change user3 data actually you change all data referenced
	to user1, because user1,user3 and user4 using same reference on memory.
	*/
	user3.UserName = "Dump"
	fmt.Println("user1 :", user1)
	fmt.Println("user4 :", user4)
	fmt.Println("user3 :", user3)
	fmt.Println("===============================")

	// Move all reference from user1 to user3
	*user3 = GithubData{"Lazycat Labs", "Mudassir", "hey.mudassir@gmail.com"}
	fmt.Println("user1 :", user1) // user1 change reference to user3
	fmt.Println("user4 :", user4) // user4 change reference to user3
	fmt.Println("user3 :", user3)

	// create new pointer
	user5 := &GithubData{"Lazycat Labs", "Mudassir", "hey.mudassir@gmail.com"}
	// create new empty pointer
	user6 := new(GithubData)

	fmt.Println("===============================")
	fmt.Println("user5 :", user5)
	fmt.Println("user6 :", user6)

	// Change pointer struct to struct
	var user5Struct GithubData = *user5
	var user6Struct GithubData = *user6

	fmt.Println("===============================")
	fmt.Println("userStruct5 :", user5Struct)
	fmt.Println("userStruct6 :", user6Struct)

	// Pointer in function
	user7 := GithubData{"Lzyct", "Mudassir", ""}
	fmt.Println("===============================")
	fmt.Println("user7 :", user7)
	updateEmail(&user7)
	fmt.Println("user7 :", user7)
}

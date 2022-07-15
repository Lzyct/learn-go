package main

import (
	"fmt"
	"sort"
)

type User struct {
	UserName string
	Age      int
}

type UserSlice []User

func (user UserSlice) Len() int {
	return len(user)
}

func (user UserSlice) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

func (user UserSlice) Less(i, j int) bool {
	// sort by age
	return user[i].Age <= user[j].Age
}

func main() {
	users := []User{
		{UserName: "Budi", Age: 30},
		{UserName: "Bude", Age: 30},
		{UserName: "Anton", Age: 30},
		{UserName: "Andi", Age: 25},
		{UserName: "Abdul", Age: 25},
		{UserName: "Abdullah", Age: 25},
	}

	// Before sort
	fmt.Println(users)

	// After sort
	sort.Sort(UserSlice(users))
	fmt.Println(users)

	// Reverse sort
	sort.Sort(sort.Reverse(UserSlice(users)))
	fmt.Println(users)
}

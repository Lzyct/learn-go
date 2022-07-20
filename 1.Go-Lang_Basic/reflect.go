package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type GithubUser struct {
	UserName string `mandatory:"true" max:"5"`
	Name     string
	Stars    int
}

/**
Create validation using refrect by tag in stuct
*/

func IsValid(data any) bool {
	t := reflect.TypeOf(data)
	// check data validation by tag
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("mandatory") == "true" && field.Tag.Get("max") != "" {
			fieldValue := reflect.ValueOf(data).Field(i).Interface()
			max, _ := strconv.Atoi(field.Tag.Get("max"))
			return fieldValue != "" && len(fieldValue.(string)) <= max
		}
	}

	return true
}

func main() {
	githubUser := GithubUser{
		UserName: "Lzyct Labs",
		Name:     "Mudassir",
		Stars:    10}

	githubUserType := reflect.TypeOf(githubUser)

	// fmt.Println(githubUserType.NumField()) // 3

	// Get all field name from struct
	for i := 0; i < githubUserType.NumField(); i++ {
		fmt.Print(githubUserType.Field(i).Name, " -> ") // Get field name
		// fmt.Println(githubUserType.Field(i).Tag.Get("mandatory"))     // Get field tag
		// fmt.Println(githubUserType.Field(i).Tag.Get("max"))           // Get field tag
		fmt.Println(reflect.ValueOf(githubUser).Field(i).Interface()) // Get field value
	}

	fmt.Println("isValid? ", IsValid(githubUser))
}

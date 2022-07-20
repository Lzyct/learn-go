package main

import "fmt"

type GithubInterface interface {
	GetUserName() string
}

func showName(githubInterface GithubInterface) {
	fmt.Println("Hello", githubInterface.GetUserName())
}

type GithubUsername struct {
	Name string
}
type Animal struct {
	Name string
}

// Implement GithubInterface
func (githubUsername GithubUsername) GetUserName() string {
	return githubUsername.Name
}

// Implement GithubInterface
func (animal Animal) GetUserName() string {
	return animal.Name
}

func main() {
	githubUserName := GithubUsername{
		Name: "Lzyct",
	}
	showName(githubUserName)

	animal := Animal{
		Name: "Cat",
	}
	showName(animal)
}

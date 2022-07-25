package services

import (
	"errors"

	"github.com/Lzyct/learn-go/3.Go-Lang_UnitTest/entities"
	"github.com/Lzyct/learn-go/3.Go-Lang_UnitTest/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (userService UserService) Get(username string, name string) (entities.Users, error) {
	user := userService.Repository.FindUser(username, name)
	if user == nil {
		return nil, errors.New("User not found")
	} else {
		return user, nil
	}
}

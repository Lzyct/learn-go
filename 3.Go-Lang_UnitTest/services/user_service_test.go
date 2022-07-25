package services

import (
	"testing"

	"github.com/Lzyct/learn-go/3.Go-Lang_UnitTest/entities"
	"github.com/Lzyct/learn-go/3.Go-Lang_UnitTest/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repositories.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestUserService_UserNotFound(t *testing.T) {
	// repository will be mock on "FindUser" function and return nil
	// make sure number of arguments is same like in your method
	// This will be simulate the repository process
	userRepository.Mock.On("FindUser", "username", "name").Return(nil)

	result, err := userService.Get("username", "name")

	println("ERROR:", err.Error())
	// should return error because the return is nil
	assert.NotNil(t, err)
	// should return nil for the result
	assert.Nil(t, result)
}

func TestUserService_UserFound(t *testing.T) {
	//create dummy response
	users := &[]entities.User{
		{
			Id:       "1",
			UserName: "Lazy",
			Name:     "Lazy",
		},
		{
			Id:       "2",
			UserName: "LazyCat",
			Name:     "LazyCat",
		},
		{
			Id:       "3",
			UserName: "LazyCatLabs",
			Name:     "LazyCatLabs",
		},
	}

	// Mock response
	userRepository.Mock.On("FindUser", "Lazy", "Lazy").Return(entities.Users(users))

	// call service from mock repository
	result, err := userService.Get("Lazy", "Lazy")

	// the result shouldn't error
	assert.Nil(t, err)

	// the result shouldn't nil
	assert.NotNil(t, result)

	// the result should be same with users
	assert.Equal(t, result, entities.Users(users))
}

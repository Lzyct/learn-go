package repositories

import (
	"github.com/Lzyct/learn-go/3.Go-Lang_UnitTest/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (userRepositoryMock UserRepositoryMock) FindUser(username string, name string) entities.Users {
	// Collect all parameters
	arguments := userRepositoryMock.Mock.Called(username, name)
	/// get arguments from index 0
	if arguments.Get(0) == nil {
		return nil
	}
	user := arguments.Get(0).(entities.Users)
	return user
}

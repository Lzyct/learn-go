package repositories

import "github.com/Lzyct/learn-go/3.Go-Lang_UnitTest/entities"

// Create contract for User Repository
type UserRepository interface {
	FindUser(username string, name string) entities.Users
}

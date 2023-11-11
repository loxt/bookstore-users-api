package services

import (
	"github.com/loxt/bookstore-users-api/domain/users"
	"github.com/loxt/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil
}

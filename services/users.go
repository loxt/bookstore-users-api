package services

import (
	"github.com/loxt/bookstore-users-api/domain/users"
	"github.com/loxt/bookstore-users-api/utils/date_utils"
	"github.com/loxt/bookstore-users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		ID: userId,
	}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{
		ID: userId,
	}

	return user.Delete()
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)

	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(&user); err != nil {
		return nil, err
	}

	return current, nil
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}

	return dao.Search(status)
}

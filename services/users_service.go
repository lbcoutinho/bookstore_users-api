package services

import (
	"github.com/lbcoutinho/bookstore_users-api/domain/users"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
)

func Create(user users.User) (*users.User, *errors.RestErr) {
	user.TrimSpace()

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func Get(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func Update(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	user.TrimSpace()

	current, err := Get(user.Id)
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

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func Delete(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

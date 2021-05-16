package services

import (
	"github.com/lbcoutinho/bookstore_users-api/domain/users"
	"github.com/lbcoutinho/bookstore_users-api/utils/crypto_utils"
	"github.com/lbcoutinho/bookstore_users-api/utils/date_utils"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
)

var UserService userServiceInterface = &userService{}

type userService struct {
}

type userServiceInterface interface {
	Create(users.User) (*users.User, *errors.RestErr)
	Get(int64) (*users.User, *errors.RestErr)
	Update(users.User, bool) (*users.User, *errors.RestErr)
	Delete(int64) *errors.RestErr
	Search(string) (users.Users, *errors.RestErr)
}

func (us *userService) Create(user users.User) (*users.User, *errors.RestErr) {
	user.TrimSpace()
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *userService) Get(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) Update(user users.User, isPartial bool) (*users.User, *errors.RestErr) {
	user.TrimSpace()

	current, err := us.Get(user.Id)
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
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (us *userService) Delete(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func (us *userService) Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

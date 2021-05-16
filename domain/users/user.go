package users

import (
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
	"strings"
)

const StatusActive = "active"

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

type Users []User

func (u *User) TrimSpace() {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
}

func (u *User) Validate() *errors.RestErr {
	u.Email = strings.ToLower(u.Email)
	if u.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}

	if u.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}

	return nil
}

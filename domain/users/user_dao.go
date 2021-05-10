package users

import (
	"github.com/lbcoutinho/bookstore_users-api/utils/date_utils"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
)

func (u *User) Get() *errors.RestErr {
	return nil
}

func (u *User) Save() *errors.RestErr {
	u.DateCreated = date_utils.GetNowString()
	return nil
}

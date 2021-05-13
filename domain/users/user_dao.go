package users

import (
	"github.com/lbcoutinho/bookstore_users-api/datasources/mysql/users_db"
	"github.com/lbcoutinho/bookstore_users-api/utils/date_utils"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
)

func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	return nil
}

func (u *User) Save() *errors.RestErr {
	u.DateCreated = date_utils.GetNowString()
	return nil
}

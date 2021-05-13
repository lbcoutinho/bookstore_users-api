package users

import (
	"fmt"
	"github.com/lbcoutinho/bookstore_users-api/datasources/mysql/users_db"
	"github.com/lbcoutinho/bookstore_users-api/utils/date_utils"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	indexUniqueEmail = "email_UNIQUE"
)

func (u *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	u.DateCreated = date_utils.GetNowString()
	result, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", u.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error while trying to save the users: %s", err.Error()))
	}

	u.Id, err = result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error while trying to save the users: %s", err.Error()))
	}

	return nil
}

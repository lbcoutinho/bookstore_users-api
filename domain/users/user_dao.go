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
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", u.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error while retrieving user id %d: %s", u.Id, err.Error()))
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

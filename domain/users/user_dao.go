package users

import (
	"github.com/lbcoutinho/bookstore_users-api/datasources/mysql/users_db"
	"github.com/lbcoutinho/bookstore_users-api/utils/date_utils"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
	"github.com/lbcoutinho/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser      = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	queryGetUserById     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
	queryUpdateUser      = "UPDATE users SET first_name=?, last_name=?, email=? where id = ?"
	errorDuplicatedEntry = 1062
	errorNoRows          = "no rows in result set"
)

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	u.DateCreated = date_utils.GetNowString()
	result, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	u.Id, err = result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

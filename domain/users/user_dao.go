package users

import (
	"github.com/lbcoutinho/bookstore_users-api/datasources/mysql/users_db"
	"github.com/lbcoutinho/bookstore_users-api/logger"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUserById      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, status=? where id = ?;"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
	errorDuplicatedEntry  = 1062
	errorNoRows           = "no rows in result set"
)

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error while trying to prepare save user statement", err)
		return newDatabaseError()
	}
	defer stmt.Close()

	result, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
	if saveErr != nil {
		logger.Error("Error while trying to save new user", saveErr)
		return newDatabaseError()
	}

	u.Id, err = result.LastInsertId()
	if err != nil {
		logger.Error("Error while trying to get last insert id after creating a new user", err)
		return newDatabaseError()
	}

	return nil
}

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		logger.Error("Error while trying to prepare get user statement", err)
		return newDatabaseError()
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); getErr != nil {
		logger.Error("Error while trying to get user by id", getErr)
		return newDatabaseError()
	}

	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error while trying to prepare update user statement", err)
		return newDatabaseError()
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id, u.Status)
	if err != nil {
		logger.Error("Error while trying to update user", err)
		return newDatabaseError()
	}

	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error while trying to prepare delete user statement", err)
		return newDatabaseError()
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		logger.Error("Error while trying to delete user", err)
		return newDatabaseError()
	}

	return nil
}

func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error while trying to prepare query users by status statement", err)
		return nil, newDatabaseError()
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error while trying to query users by status", err)
		return nil, newDatabaseError()
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error while trying to scan user row into user struct", err)
			return nil, newDatabaseError()
		}

		results = append(results, user)
	}

	return results, nil
}

func newDatabaseError() *errors.RestErr {
	return errors.NewInternalServerError("Database error")
}

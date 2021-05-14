package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/lbcoutinho/bookstore_users-api/utils/errors"
	"strings"
)

const (
	errorDuplicatedEntry = 1062
	errorNoRows          = "no rows in result set"
)


func ParseError(err error) *errors.RestErr {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}

	if mysqlErr.Number == errorDuplicatedEntry {
		return errors.NewBadRequestError("Invalid data. Unique index violated.")
	}
	return errors.NewInternalServerError("Error processing request")
}
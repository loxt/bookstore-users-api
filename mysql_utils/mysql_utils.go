package mysql_utils

import (
	e "errors"
	"github.com/go-sql-driver/mysql"
	"github.com/loxt/bookstore-users-api/utils/errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	var sqlErr *mysql.MySQLError
	ok := e.As(err, &sqlErr)

	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}

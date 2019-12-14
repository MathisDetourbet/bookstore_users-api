package errors

import (
	"log"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRow = "no rows in result set"
)

func ParseError(err error) *RestErr {
	sqlErr, isSQLErr := err.(*mysql.MySQLError)
	if isSQLErr {
		switch sqlErr.Number {
		case 1062:
			return NewBadRequestError("invalid data")
		}
		return NewInternalServerError("error processing request")
	}

	if strings.Contains(err.Error(), errorNoRow) {
		return NewNotFoundError("no record matching the given id")
	}
	log.Println(err.Error())
	return NewInternalServerError("error when parsing database response")
}

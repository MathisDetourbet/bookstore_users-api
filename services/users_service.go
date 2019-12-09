package services

import (
	"github.com/MathisDetourbet/bookstore_users-api/app/utils/errors"
	"github.com/MathisDetourbet/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

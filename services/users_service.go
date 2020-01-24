package services

import (
	"github.com/MathisDetourbet/bookstore_users-api/domain/users"
	"github.com/MathisDetourbet/bookstore_users-api/utils/crypto"
	"github.com/MathisDetourbet/bookstore_users-api/utils/date"
	"github.com/MathisDetourbet/bookstore_users-api/utils/errors"
)

// GetUser contains the logic to get a user from the database by using its ID
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser contains the logic to create a new user: validate input data and save these data in the database
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBFormat()
	user.DateUpdated = user.DateCreated
	user.Password = crypto.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser contains the logic to update a user: get the existing user then (if it exists) update it with the new values
func UpdateUser(isPartial bool, user *users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
		if user.Password != "" {
			currentUser.Password = user.Password
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
		currentUser.Password = user.Password
	}

	currentUser.DateUpdated = date.GetNowDBFormat()
	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

// DeleteUser contains the logic of deleting a single user in the database by giving an user ID
func DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

// Search function contains the logic of finding a user by its status field
func Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

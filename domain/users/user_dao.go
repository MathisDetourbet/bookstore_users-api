package users

import (
	"fmt"

	"github.com/MathisDetourbet/bookstore_users-api/datasources/mysql/users_db"
	"github.com/MathisDetourbet/bookstore_users-api/utils/date"
	"github.com/MathisDetourbet/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, date_updated) VALUES(?, ?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created, date_updated FROM users WHERE id=?"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=?, date_updated=? WHERE id=?"
)

// Get a user from the database
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	getResult := stmt.QueryRow(user.ID)
	if getErr := getResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.DateUpdated); getErr != nil {
		return errors.ParseError(getErr)
	}
	return nil
}

// Save a user in the database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	user.DateUpdated = user.DateCreated

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.DateUpdated)
	if saveErr != nil {
		return errors.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.ID = userID
	return nil
}

// Update a user previously recorded in the database
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateUpdated = date.GetNowString()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateUpdated, user.ID)
	if updateErr != nil {
		return errors.ParseError(updateErr)
	}
	if err := user.Get(); err != nil {
		return err
	}

	return nil
}

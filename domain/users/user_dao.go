package users

import (
	"fmt"
	"log"

	"github.com/MathisDetourbet/bookstore_users-api/datasources/mysql/users_db"
	"github.com/MathisDetourbet/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, date_updated, status, password) VALUES(?, ?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, date_updated, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, date_updated=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, date_updated, status FROM users WHERE status=?;"
)

// Get a user from the database
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	getResult := stmt.QueryRow(user.ID)
	if getErr := getResult.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.DateUpdated, &user.Status); getErr != nil {
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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.DateUpdated, user.Status, user.Password)
	if saveErr != nil {
		log.Println(saveErr.Error())
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

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateUpdated, user.ID)
	if updateErr != nil {
		return errors.ParseError(updateErr)
	}
	if err := user.Get(); err != nil {
		return err
	}

	return nil
}

// Delete a user in the database by its id
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, deleteErr := stmt.Exec(user.ID); deleteErr != nil {
		return errors.ParseError(deleteErr)
	}
	return nil
}

// FindByStatus is finding a user by its status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, findErr := stmt.Query(status)
	if findErr != nil {
		return nil, errors.ParseError(findErr)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.DateUpdated, &user.Status); err != nil {
			return nil, errors.ParseError(err)
		}
		results = append(results, user)
	}
	return results, nil
}

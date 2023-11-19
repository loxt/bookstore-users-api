package users

import (
	"fmt"
	"github.com/loxt/bookstore-users-api/datasources/mysql/users_db"
	"github.com/loxt/bookstore-users-api/logger"
	"github.com/loxt/bookstore-users-api/mysql_utils"
	"github.com/loxt/bookstore-users-api/utils/errors"
	"strings"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, password, status, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users where email = ? AND password = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.Status, user.DateCreated)

	if err != nil {
		logger.Error("error when trying to save user", err)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}

	user.ID = userId

	return nil
}

func (user *User) Update(newUser *User) *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.FirstName, newUser.LastName, newUser.Email, user.ID)

	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
func (user *User) Search(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)

	if err != nil {
		logger.Error("error when trying to prepare search users statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		logger.Error("error when trying to search users", err)
		return nil, errors.NewInternalServerError("database error")
	}

	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)

	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

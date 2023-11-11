package users

import (
	"github.com/loxt/bookstore-users-api/datasources/mysql/users_db"
	"github.com/loxt/bookstore-users-api/mysql_utils"
	"github.com/loxt/bookstore-users-api/utils/date_utils"
	"github.com/loxt/bookstore-users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userId

	return nil
}

func (user *User) Update(newUser *User) *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(newUser.FirstName, newUser.LastName, newUser.Email, user.ID)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

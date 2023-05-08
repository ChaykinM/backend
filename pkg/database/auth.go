package database

import (
	"fmt"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) LoginAuthorization(loginRequest *models.LoginRequest) (models.AuthUserData, error) {
	request := fmt.Sprintf("SELECT id, first_name, second_name, email, status FROM public.users WHERE login='%s' AND password='%s';", loginRequest.Login, loginRequest.Password)

	row := d.dbDriver.QueryRow(request)
	var id int
	var first_name, second_name, email, status string
	err := row.Scan(&id, &first_name, &second_name, &email, &status)
	var authData models.AuthUserData
	authData.UserID = id
	authData.FirstName = first_name
	authData.SecondName = second_name
	authData.Email = email
	authData.Status = status
	return authData, err
}

func (d *Database) RegisterUser(registerRequest *models.RegisterRequest) (models.AuthUserData, error) {
	var authData models.AuthUserData

	request := fmt.Sprintf("INSERT INTO public.users(login, first_name, second_name, email, password, status) VALUES('%s', '%s', '%s', '%s', '%s', 'employee') RETURNING id;", registerRequest.Login, registerRequest.FirstName, registerRequest.SecondName, registerRequest.Email, registerRequest.Password)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return authData, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int

	err = row.Scan(&id)
	authData.UserID = id
	authData.FirstName = registerRequest.FirstName
	authData.SecondName = registerRequest.SecondName
	authData.Email = registerRequest.Email
	authData.Status = "employee"

	return authData, err
}

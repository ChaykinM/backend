package database

import (
	"fmt"

	"main.go/pkg/models"
	// "../models"
)

func (d *Database) GetUsers() ([]*models.User, error) {
	var users []*models.User
	request := fmt.Sprintf("SELECT id, reg_time, login, email, first_name, second_name, status FROM public.users")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return users, err
	} else {
		for rows.Next() {
			var id int
			var reg_time, login, email, first_name, second_name, status string
			rows.Scan(&id, &reg_time, &login, &email, &first_name, &second_name, &status)
			var user models.User
			user.ID = id
			user.RegTime = reg_time
			user.Login = login
			user.Email = email
			user.FirstName = first_name
			user.SecondName = second_name
			user.Status = status
			users = append(users, &user)
		}
	}
	return users, nil

}

func (d *Database) GetUserByID(id int) (*models.User, error) {
	var user models.User
	request := fmt.Sprintf("SELECT reg_time, login, email, first_name, second_name, status FROM public.users WHERE id = %d;", id)
	rows := d.dbDriver.QueryRow(request)

	var reg_time, login, email, first_name, second_name, status string
	err := rows.Scan(&reg_time, &login, &email, &first_name, &second_name, &status)
	user.ID = id
	user.RegTime = reg_time
	user.Login = login
	user.Email = email
	user.FirstName = first_name
	user.SecondName = second_name
	user.Status = status

	return &user, err
}

func (d *Database) EditUser(userUpdate *models.UserUpdate) error {
	request := fmt.Sprintf("UPDATE public.users SET login = '%s', email = '%s', status = '%s' WHERE id = %d", userUpdate.Login, userUpdate.Email, userUpdate.Status, userUpdate.ID)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}

func (d *Database) DeleteUser(id int) error {
	request := fmt.Sprintf("DELETE FROM public.users WHERE id = %d", id)
	_, err := d.dbDriver.Exec(request)
	return err
}

func (d *Database) UpdateUserPassword(userUpdPass *models.UserUpdPassRequest) error {
	request := fmt.Sprintf("UPDATE public.users SET password = '%s' WHERE id = %d", userUpdPass.Password, userUpdPass.ID)

	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	return err
}

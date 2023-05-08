package models

type User struct {
	ID         int    `json:"id"`
	RegTime    string `json:"reg_time"`
	Login      string `json:"login"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Status     string `json:"status"`
}

type UserUpdate struct {
	ID     int    `json:"id"`
	Login  string `json:"login"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type UserUpdPassRequest struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

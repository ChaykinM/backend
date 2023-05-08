package models

// func (d *Database)

type TasksCount struct {
	FirstLevel  int `json:"first_level"`
	SecondLevel int `json:"second_level"`
	ThirdLevel  int `json:"third_level"`
}

type UserRating struct {
	FirstLevel  int `json:"first_level"`
	SecondLevel int `json:"second_level"`
	ThirdLevel  int `json:"third_level"`
}

type Ratings struct {
	FirstLevel  []RatingTableRow `json:"first_level"`
	SecondLevel []RatingTableRow `json:"second_level"`
	ThirdLevel  []RatingTableRow `json:"third_level"`
}

type RatingTableRow struct {
	UserID int    `json:"user_id"`
	Login  string `json:"login"`
	Count  int    `json:"count"`
}

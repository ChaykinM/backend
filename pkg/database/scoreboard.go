package database

import (
	"context"
	"fmt"

	"main.go/pkg/models"
)

func (d *Database) GetTasksCount() (*models.TasksCount, error) {
	var tasksCount models.TasksCount

	request := "SELECT(SELECT COUNT(*) FROM public.tasks WHERE level = 1) AS first_level, (SELECT COUNT(*) FROM public.tasks WHERE level = 2) AS second_level, (SELECT COUNT(*) FROM public.tasks WHERE level = 3) AS third_level;"
	rows := d.dbDriver.QueryRow(request)

	var firstLevel, secondLevel, thirdLevel int
	err := rows.Scan(&firstLevel, &secondLevel, &thirdLevel)

	tasksCount.FirstLevel = firstLevel
	tasksCount.SecondLevel = secondLevel
	tasksCount.ThirdLevel = thirdLevel

	return &tasksCount, err
}

func (d *Database) GetUserRating(user_id int) (*models.UserRating, error) {
	var userRating models.UserRating

	request := fmt.Sprintf("SELECT(SELECT COUNT(public.solved_tasks.task_id) FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id WHERE public.solved_tasks.user_id = %d AND public.tasks.level = 1) AS first_level, (SELECT COUNT(public.solved_tasks.task_id) FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id WHERE public.solved_tasks.user_id = %d AND public.tasks.level = 2) AS second_level, (SELECT COUNT(public.solved_tasks.task_id) FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id WHERE public.solved_tasks.user_id = %d AND public.tasks.level = 3) AS third_level;", user_id, user_id, user_id)
	rows := d.dbDriver.QueryRow(request)

	var firstLevel, secondLevel, thirdLevel int
	err := rows.Scan(&firstLevel, &secondLevel, &thirdLevel)

	userRating.FirstLevel = firstLevel
	userRating.SecondLevel = secondLevel
	userRating.ThirdLevel = thirdLevel

	return &userRating, err
}

func (d *Database) GetRatings() (*models.Ratings, error) {
	var ratings models.Ratings

	ctx := context.Background()
	tx, err := d.dbDriver.BeginTx(ctx, nil)
	if err != nil {
		return &ratings, err
	}

	firstLevel_request := "SELECT public.solved_tasks.user_id, public.users.login, COUNT(public.solved_tasks.task_id) FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id JOIN public.users ON public.solved_tasks.user_id = public.users.id WHERE public.tasks.level = 1 GROUP BY public.solved_tasks.user_id, public.users.login ORDER BY count DESC; "
	rows, err := d.dbDriver.QueryContext(ctx, firstLevel_request)
	if err != nil {
		return &ratings, err
	} else {
		var ratingRows []models.RatingTableRow
		for rows.Next() {
			var user_id, count int
			var login string
			rows.Scan(&user_id, &login, &count)
			var ratingRow models.RatingTableRow
			ratingRow.UserID = user_id
			ratingRow.Login = login
			ratingRow.Count = count
			ratingRows = append(ratingRows, ratingRow)
		}
		ratings.FirstLevel = ratingRows
	}

	secondLevel_request := "SELECT public.solved_tasks.user_id, public.users.login, COUNT(public.solved_tasks.task_id) FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id JOIN public.users ON public.solved_tasks.user_id = public.users.id WHERE public.tasks.level = 2 GROUP BY public.solved_tasks.user_id, public.users.login ORDER BY count DESC; "
	rows, err = d.dbDriver.QueryContext(ctx, secondLevel_request)
	if err != nil {
		return &ratings, err
	} else {
		var ratingRows []models.RatingTableRow
		for rows.Next() {
			var user_id, count int
			var login string
			rows.Scan(&user_id, &login, &count)
			var ratingRow models.RatingTableRow
			ratingRow.UserID = user_id
			ratingRow.Login = login
			ratingRow.Count = count
			ratingRows = append(ratingRows, ratingRow)
		}
		ratings.SecondLevel = ratingRows
	}

	thirdLevel_request := "SELECT public.solved_tasks.user_id, public.users.login, COUNT(public.solved_tasks.task_id) FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id JOIN public.users ON public.solved_tasks.user_id = public.users.id WHERE public.tasks.level = 3 GROUP BY public.solved_tasks.user_id, public.users.login ORDER BY count DESC; "
	rows, err = d.dbDriver.QueryContext(ctx, thirdLevel_request)
	if err != nil {
		return &ratings, err
	} else {
		var ratingRows []models.RatingTableRow
		for rows.Next() {
			var user_id, count int
			var login string
			rows.Scan(&user_id, &login, &count)
			var ratingRow models.RatingTableRow
			ratingRow.UserID = user_id
			ratingRow.Login = login
			ratingRow.Count = count
			ratingRows = append(ratingRows, ratingRow)
		}
		ratings.ThirdLevel = ratingRows
	}

	err = tx.Commit()
	return &ratings, err
}

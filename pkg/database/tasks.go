package database

import (
	"fmt"

	"main.go/pkg/models"
)

func (d *Database) GetTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	request := fmt.Sprintf("SELECT id, create_time, title, description, level FROM public.tasks;")
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return tasks, err
	} else {
		for rows.Next() {
			var id, level int
			var create_time, title, description string
			rows.Scan(&id, &create_time, &title, &description, &level)
			var task models.Task
			task.Id = id
			task.CreateTime = create_time
			task.Title = title
			task.Description = description
			task.Level = level
			tasks = append(tasks, &task)
		}
	}

	return tasks, nil
}

func (d *Database) GetUserSolvedTasks(user_id int) ([]int, error) {
	var tasks_ids []int
	request := fmt.Sprintf("SELECT public.solved_tasks.task_id FROM public.solved_tasks JOIN public.tasks ON public.solved_tasks.task_id = public.tasks.id WHERE public.solved_tasks.user_id = %d;", user_id)
	rows, err := d.dbDriver.Query(request)
	if err != nil {
		return tasks_ids, err
	} else {
		for rows.Next() {
			var id int
			rows.Scan(&id)
			tasks_ids = append(tasks_ids, id)
		}
	}

	return tasks_ids, nil
}

func (d *Database) AddTask(task *models.AddTaskRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.tasks(title, description, level) VALUES('%s', '%s', '%d') RETURNING id;", task.Title, task.Description, task.Level)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}

func (d *Database) SolveTask(solvedTask *models.SolveTaskRequest) (int, error) {
	request := fmt.Sprintf("INSERT INTO public.solved_tasks(task_id, user_id) VALUES('%d', '%d') RETURNING id;", solvedTask.TaskId, solvedTask.UserId)
	stmt, err := d.dbDriver.Prepare(request)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int
	err = row.Scan(&id)
	return id, err
}

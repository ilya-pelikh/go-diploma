package task

import (
	"database/sql"
	"errors"
	"fmt"

	"diploma/internal/pkg/constants"
	"diploma/internal/pkg/db"
)

type repository struct{}

var Repository repository

func (r *repository) AddTask(dto *AddTaskRequestDTO) (*AddTaskResponseDTO, error) {
	result, err := db.Database.Exec(`
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (:date, :title, :comment, :repeat)`,
		sql.Named("date", dto.Date),
		sql.Named("title", dto.Title),
		sql.Named("comment", dto.Comment),
		sql.Named("repeat", dto.Repeat))

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	resp := AddTaskResponseDTO{
		Id:      id,
		date:    dto.Date,
		title:   dto.Title,
		comment: dto.Comment,
		repeat:  dto.Repeat,
	}

	return &resp, nil
}

func (r *repository) GetAllTasks(search string, date string) ([]*TaskResponseDTO, error) {
	query := `
		SELECT id, date, title, comment, repeat
		FROM scheduler`

	if len(date) != 0 {
		query += `
		WHERE date = :date`
	} else {
		query += `
		WHERE 1=1`
	}

	if len(search) != 0 {
		query += `
		AND title LIKE :search
		OR comment LIKE :search`
	}

	query += fmt.Sprintf(`
		ORDER BY date
		LIMIT %d`, constants.PageSize)

	rows, err := db.Database.Query(query,
		sql.Named("search", "%"+search+"%"),
		sql.Named("date", date),
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*TaskResponseDTO, 0)

	for rows.Next() {
		task := TaskResponseDTO{}

		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) GetTaskById(id string) (*TaskResponseDTO, error) {
	row := db.Database.QueryRow(`
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = :id`,
		sql.Named("id", id))

	task := TaskResponseDTO{}

	err := row.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *repository) UpdateTask(dto *UpdateTaskRequestDTO) error {
	result, err := db.Database.Exec(`
		UPDATE scheduler
		SET date = :date,
			title = :title,
			comment = :comment,
			repeat = :repeat
		WHERE id = :id`,
		sql.Named("date", dto.Date),
		sql.Named("title", dto.Title),
		sql.Named("comment", dto.Comment),
		sql.Named("repeat", dto.Repeat),
		sql.Named("id", dto.Id),
	)

	if err != nil {
		return err
	}

	affectedRowsCount, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affectedRowsCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *repository) DeleteTask(id string) error {
	result, err := db.Database.Exec(`
		DELETE FROM scheduler
		WHERE id = :id`,
		sql.Named("id", id),
	)

	if err != nil {
		return err
	}

	affectedRowsCount, err := result.RowsAffected()

	if err != nil {
		return err
	}
	if affectedRowsCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

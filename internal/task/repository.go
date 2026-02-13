package task

import (
	"database/sql"
	"diploma/internal/pkg/logger"

	"go.uber.org/zap"
)

type repository struct {
	db string
}

var Repository repository

func (r *repository) AddTask(path string, dto *AddTaskRequestDTO) (*AddTaskResponseDTO, error) {
	db, err := sql.Open("sqlite", path)

	if err != nil {
		logger.Get().Error("SQL database file not found")
		return nil, err
	}
	defer db.Close()

	result, err := db.Exec(`
		INSERT INTO scheduler (date, title, comment, repeat)
		VALUES (:date, :title, :comment, :repeat)`,
		sql.Named("date", dto.Date),
		sql.Named("title", dto.Title),
		sql.Named("comment", dto.Comment),
		sql.Named("repeat", dto.Repeat))

	if err != nil {
		logger.Get().Error("Couldn't add task using sql", zap.Error(err))
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Get().Error("Couldn't get task id from sql result", zap.Error(err))
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

func (r *repository) GetAllTasks(path string, search string, date string) ([]*TaskResponseDTO, error) {
	db, err := sql.Open("sqlite", path)

	if err != nil {
		logger.Get().Error("SQL database file not found")
		return nil, err
	}

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

	query += `
		ORDER BY date
		LIMIT 50`

	rows, err := db.Query(query,
		sql.Named("search", "%"+search+"%"),
		sql.Named("date", date),
	)

	if err != nil {
		logger.Get().Error("Couldn't get all task using sql")
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*TaskResponseDTO, 0)

	for rows.Next() {
		task := TaskResponseDTO{}

		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			logger.Get().Error("Couldn't parse task after using sql")
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		logger.Get().Error("Rows error")
		return nil, err
	}

	return tasks, nil
}

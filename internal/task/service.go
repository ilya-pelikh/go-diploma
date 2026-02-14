package task

import (
	"diploma/internal/pkg/env"
	"diploma/internal/planner"
	"strings"
	"time"
)

func AddTask(dto *AddTaskRequestDTO) (*AddTaskResponseDTO, error) {
	if strings.TrimSpace(dto.Date) == "" {
		dto.Date = time.Now().Format("20060102")
	} else {

		if dto.Date < time.Now().Format("20060102") {
			if strings.TrimSpace(dto.Repeat) == "" {
				dto.Date = time.Now().Format("20060102")
			} else {
				date, err := planner.NextDate(time.Now(), dto.Date, dto.Repeat)
				if err != nil {
					return nil, err
				}

				dto.Date = date
			}
		}

	}

	result, err := Repository.AddTask(env.TODO_DBFILE, dto)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllTasks(search string, date string) ([]*TaskResponseDTO, error) {
	result, err := Repository.GetAllTasks(env.TODO_DBFILE, search, date)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTaskById(id string) (*TaskResponseDTO, error) {
	result, err := Repository.GetTaskById(env.TODO_DBFILE, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateTask(dto *UpdateTaskRequestDTO) error {
	if strings.TrimSpace(dto.Date) == "" {
		dto.Date = time.Now().Format("20060102")
	} else {

		if dto.Date < time.Now().Format("20060102") {
			if strings.TrimSpace(dto.Repeat) == "" {
				dto.Date = time.Now().Format("20060102")
			} else {
				date, err := planner.NextDate(time.Now(), dto.Date, dto.Repeat)
				if err != nil {
					return err
				}

				dto.Date = date
			}
		}

	}
	err := Repository.UpdateTask(env.TODO_DBFILE, dto)

	if err != nil {
		return err
	}
	return nil
}

func DoTask(id string) error {
	result, err := Repository.GetTaskById(env.TODO_DBFILE, id)
	if err != nil {
		return err
	}

	if result.Repeat == "" {
		err := Repository.DeleteTask(env.TODO_DBFILE, id)
		if err != nil {
			return err
		}
		return nil
	}

	date, err := planner.NextDate(time.Now(), result.Date, result.Repeat)
	if err != nil {
		return err
	}

	task := &UpdateTaskRequestDTO{
		Id:      result.Id,
		Date:    date,
		Title:   result.Title,
		Comment: result.Comment,
		Repeat:  result.Repeat,
	}

	err = Repository.UpdateTask(env.TODO_DBFILE, task)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	err := Repository.DeleteTask(env.TODO_DBFILE, id)
	if err != nil {
		return err
	}

	return nil
}

package task

import (
	"strings"
	"time"

	"diploma/internal/pkg/constants"
	"diploma/internal/planner"
)

func AddTask(dto *AddTaskRequestDTO) (*AddTaskResponseDTO, error) {
	if strings.TrimSpace(dto.Date) == "" {
		dto.Date = time.Now().Format(constants.DateFormat)
	} else {

		if dto.Date < time.Now().Format(constants.DateFormat) {
			if strings.TrimSpace(dto.Repeat) == "" {
				dto.Date = time.Now().Format(constants.DateFormat)
			} else {
				date, err := planner.NextDate(time.Now(), dto.Date, dto.Repeat)
				if err != nil {
					return nil, err
				}

				dto.Date = date
			}
		}

	}

	result, err := Repository.AddTask(dto)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllTasks(search string, date string) ([]*TaskResponseDTO, error) {
	result, err := Repository.GetAllTasks(search, date)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTaskById(id string) (*TaskResponseDTO, error) {
	result, err := Repository.GetTaskById(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateTask(dto *UpdateTaskRequestDTO) error {
	if strings.TrimSpace(dto.Date) == "" {
		dto.Date = time.Now().Format(constants.DateFormat)
	} else {

		if dto.Date < time.Now().Format(constants.DateFormat) {
			if strings.TrimSpace(dto.Repeat) == "" {
				dto.Date = time.Now().Format(constants.DateFormat)
			} else {
				date, err := planner.NextDate(time.Now(), dto.Date, dto.Repeat)
				if err != nil {
					return err
				}

				dto.Date = date
			}
		}

	}
	err := Repository.UpdateTask(dto)

	if err != nil {
		return err
	}
	return nil
}

func DoTask(id string) error {
	result, err := Repository.GetTaskById(id)
	if err != nil {
		return err
	}

	if result.Repeat == "" {
		err := Repository.DeleteTask(id)
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

	err = Repository.UpdateTask(task)

	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	err := Repository.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

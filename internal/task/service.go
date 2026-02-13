package task

import (
	"diploma/internal/pkg/env"
	"diploma/internal/planner"
	"net/http"
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

func GetAllTasks(search string, date string) ([]*TaskResponseDTO, error, int) {
	result, err := Repository.GetAllTasks(env.TODO_DBFILE, search, date)

	if err != nil {
		return nil, err, http.StatusServiceUnavailable
	}

	return result, nil, http.StatusOK
}

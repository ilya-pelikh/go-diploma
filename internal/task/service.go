package task

import (
	"diploma/internal/pkg/env"
	"net/http"
	"strings"
	"time"
)

func AddTask(dto *AddTaskRequestDTO) (*AddTaskResponsetDTO, error, int) {
	if strings.TrimSpace(dto.Date) == "" {
		dto.Date = time.Now().Format("20060102")
	} else {
		// COMMENT: я сделал проверку на уровне хендлера, решил что второй раз не нужно это делать
		data, _ := time.Parse("20060102", dto.Date)
		if data.Before(time.Now()) {
			if strings.TrimSpace(dto.Repeat) == "" {
				dto.Date = time.Now().Format("20060102")
			} else {
				// nextDay
			}
		}

	}

	result, err := Repository.AddTask(env.TODO_DBFILE, dto)

	if err != nil {
		return nil, err, http.StatusServiceUnavailable
	}

	return result, nil, http.StatusCreated
}

func GetAllTasks() ([]*AddTaskResponsetDTO, error, int) {
	result, err := Repository.GetAllTasks(env.TODO_DBFILE)

	if err != nil {
		return nil, err, http.StatusServiceUnavailable
	}

	return result, nil, http.StatusOK
}

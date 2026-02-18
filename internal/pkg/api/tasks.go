package api

import (
	"net/http"
	"time"

	"diploma/internal/pkg/constants"
	"diploma/internal/pkg/logger"
	"diploma/internal/task"
)

type tasksResponse struct {
	Tasks []*task.TaskResponseDTO `json:"tasks"`
}

func HandleTasks(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		res.Header().Set("Content-Type", "application/json; charset=utf-8")

		query := req.URL.Query()
		search := query.Get("search")

		date := ""
		text := ""

		parsedDate, err := time.Parse("02.01.2006", search)
		if err == nil {
			date = parsedDate.Format(constants.DateFormat)
		} else {
			text = search
		}

		response, err := task.GetAllTasks(text, date)

		if err != nil {
			http.Error(res, err.Error(), http.StatusServiceUnavailable)
			return
		}

		writeResponse(res, tasksResponse{Tasks: response}, http.StatusOK)

		logger.Get().Info("Sent tasks")
		return
	}

	writeResponse(res, nil, http.StatusMethodNotAllowed)
}

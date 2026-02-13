package api

import (
	"encoding/json"
	"net/http"
	"time"

	"diploma/internal/pkg/logger"
	"diploma/internal/task"
)

type tasksResponse struct {
	Tasks []*task.TaskResponseDTO `json:"tasks"`
}

func handleTasks(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		query := req.URL.Query()
		search := query.Get("search")

		date := ""
		text := ""

		parsedDate, err := time.Parse("02.01.2006", search)
		if err == nil {
			date = parsedDate.Format("20060102")
		} else {
			text = search
		}

		tasks, err := task.GetAllTasks(text, date)

		if err != nil {
			http.Error(res, err.Error(), http.StatusServiceUnavailable)
			return
		}

		resp, err := json.Marshal(tasksResponse{Tasks: tasks})
		if err != nil {
			http.Error(res, err.Error(), http.StatusServiceUnavailable)
			return
		}

		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.Write(resp)

		logger.Get().Info("Sent tasks")
	default:
		res.WriteHeader(405)
	}

}

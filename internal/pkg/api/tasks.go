package api

import (
	"encoding/json"
	"net/http"

	"diploma/internal/pkg/logger"
	"diploma/internal/task"
)

type tasksResponse struct {
	Tasks []*task.AddTaskResponsetDTO `json:"tasks"`
}

func handleTasks(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		tasks, err, status := task.GetAllTasks()

		if err != nil {
			http.Error(res, err.Error(), status)
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

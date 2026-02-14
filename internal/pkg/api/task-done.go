package api

import (
	"diploma/internal/task"
	"fmt"
	"net/http"
)

func handleTaskDone(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		res.Header().Set("Content-Type", "application/json; charset=utf-8")

		query := req.URL.Query()
		id := query.Get("id")
		if len(id) == 0 {
			writeError(res, "Нет идентификатора задачи", 400)
		}

		err := task.DoTask(id)

		if err != nil {
			writeError(res, fmt.Sprint(err), 400)
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(`{}`))
	default:
		res.WriteHeader(405)
	}
}

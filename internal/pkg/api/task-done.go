package api

import (
	"fmt"
	"net/http"

	"diploma/internal/task"
)

func HandleTaskDone(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
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
		writeResponse(res, nil, http.StatusCreated)
		return
	}

	writeResponse(res, nil, http.StatusMethodNotAllowed)
}

package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"diploma/internal/pkg/logger"
	"diploma/internal/task"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var v = validator.New()

func handleTask(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var buf bytes.Buffer
		var dto task.AddTaskRequestDTO
		if _, err := buf.ReadFrom(req.Body); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			logger.Get().Warn("Couldn't parse body:", zap.Error(err))
			return
		}

		if err := json.Unmarshal(buf.Bytes(), &dto); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			logger.Get().Warn("Couldn't parse task:", zap.Error(err))
			return
		}

		if err := v.Struct(dto); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			logger.Get().Warn("Validation error:", zap.Error(err))
			return
		}

		respose, err, status := task.AddTask(&dto)

		if err != nil {
			http.Error(res, err.Error(), status)
			logger.Get().Error("Error:", zap.Error(err))
			return
		}

		resp, err := json.Marshal(respose)
		if err != nil {
			http.Error(res, err.Error(), status)
			logger.Get().Error("Error:", zap.Error(err))
			return
		}

		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(status)
		res.Write(resp)

		logger.Get().Info("Created task:", zap.Any("task", respose))
	default:
		res.WriteHeader(405)
	}

}

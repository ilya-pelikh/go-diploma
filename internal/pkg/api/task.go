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
		res.Header().Set("Content-Type", "application/json; charset=utf-8")

		var buf bytes.Buffer
		var dto task.AddTaskRequestDTO

		if _, err := buf.ReadFrom(req.Body); err != nil {
			writeError(res, "Не указан заголовок задачи", http.StatusBadRequest)
			logger.Get().Warn("Couldn't parse body:", zap.Error(err))
			return
		}

		if err := json.Unmarshal(buf.Bytes(), &dto); err != nil {
			writeError(res, "Не указан заголовок задачи", http.StatusBadRequest)
			logger.Get().Warn("Couldn't parse task:", zap.Error(err))
			return
		}

		if err := v.Struct(dto); err != nil {
			logger.Get().Warn("Validation error:", zap.Error(err))

			if ve, ok := err.(validator.ValidationErrors); ok {
				for _, fe := range ve {
					if fe.Field() == "Date" && fe.Tag() == "datetime" {
						writeError(res, "Некорректное значение поля date", http.StatusBadRequest)
						return
					}
					if fe.Field() == "Title" && fe.Tag() == "required" {
						writeError(res, "Не указан заголовок задачи", http.StatusBadRequest)
						return
					}
				}
			}

			writeError(res, "Некорректный запрос", http.StatusBadRequest)
			return
		}

		respose, err := task.AddTask(&dto)

		if err != nil {
			writeError(res, err.Error(), http.StatusInternalServerError)
			logger.Get().Error("Error:", zap.Error(err))
			return
		}

		resp, err := json.Marshal(respose)
		if err != nil {
			writeError(res, err.Error(), http.StatusInternalServerError)
			logger.Get().Error("Error:", zap.Error(err))
			return
		}
		res.WriteHeader(http.StatusCreated)
		res.Write(resp)

		logger.Get().Info("Created task:", zap.Any("task", respose))
	default:
		res.WriteHeader(405)
	}

}

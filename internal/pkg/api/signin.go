package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"diploma/internal/identity"
	"diploma/internal/pkg/logger"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func HandleSignin(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		res.Header().Set("Content-Type", "application/json; charset=utf-8")

		var buf bytes.Buffer
		var dto identity.AuthRequestDTO

		if _, err := buf.ReadFrom(req.Body); err != nil {
			writeError(res, "Нет тела авторизации", http.StatusBadRequest)
			logger.Get().Warn("Couldn't parse body:", zap.Error(err))
			return
		}

		if err := json.Unmarshal(buf.Bytes(), &dto); err != nil {
			writeError(res, "Нет тела авторизации", http.StatusBadRequest)
			logger.Get().Warn("Couldn't parse task:", zap.Error(err))
			return
		}

		if err := v.Struct(dto); err != nil {
			logger.Get().Warn("Validation error:", zap.Error(err))

			if ve, ok := err.(validator.ValidationErrors); ok {
				for _, fe := range ve {
					if fe.Field() == "Password" && fe.Tag() == "required" {
						writeError(res, "Не указан пароль", http.StatusBadRequest)
						return
					}
				}
			}

			writeError(res, "Некорректный запрос", http.StatusBadRequest)
			return
		}

		token, err := identity.CreateToken(dto)

		if err != nil {
			writeError(res, err.Error(), http.StatusBadRequest)
			return
		}

		writeResponse(res, token, http.StatusOK)
		return
	}

	writeResponse(res, nil, http.StatusMethodNotAllowed)
}

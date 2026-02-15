package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"diploma/internal/pkg/logger"

	"go.uber.org/zap"
)

type apiError struct {
	Error string `json:"error"`
}

func writeError(res http.ResponseWriter, text string, status int) {
	error, _ := json.Marshal(&apiError{
		Error: text,
	})

	res.WriteHeader(status)
	_, err := res.Write(error)

	if err != nil {
		logger.Logger.Error("Couldn't write error: ", zap.String("text", text))
	}
}

func writeResponse(res http.ResponseWriter, data any, status int) {
	response := []byte(`{}`)
	var err error

	switch v := data.(type) {
	case nil:

	case string:
		response = []byte(v)
	case []byte:
		response = v
	default:
		response, err = json.Marshal(v)
	}

	if err != nil {
		writeError(res, err.Error(), http.StatusInternalServerError)
		logger.Get().Error("Error:", zap.Error(err))
		return
	}

	res.WriteHeader(status)
	if _, err = res.Write(response); err != nil {
		logger.Logger.Error("Couldn't write error: ", zap.String("text", fmt.Sprintf("%+v", data)))
	}
}

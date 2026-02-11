package api

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	Error string `json:"error"`
}

func writeError(res http.ResponseWriter, text string, status int) {
	error, _ := json.Marshal(&apiError{
		Error: text,
	})

	res.WriteHeader(status)
	res.Write(error)
}

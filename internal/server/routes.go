package http

import (
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("web")))
}

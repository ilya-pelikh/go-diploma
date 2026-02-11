package api

import (
	"diploma/internal/planner"
	"fmt"
	"net/http"
	"time"
)

func handlePlanner(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":

		query := req.URL.Query()

		now, err := time.Parse("20060102", query.Get("now"))
		if err != nil {
			writeError(res, "", 400)
			return
		}

		dtstart := query.Get("date")
		if len(dtstart) == 0 {
			writeError(res, "", 400)
			return
		}

		repeat := query.Get("repeat")
		if len(repeat) == 0 {
			writeError(res, "", 400)
			return
		}

		resp, err := planner.NextDate(now, dtstart, repeat)

		if err != nil {
			writeError(res, fmt.Sprint(err), 400)
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(resp))
	}
}

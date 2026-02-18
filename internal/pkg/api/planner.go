package api

import (
	"fmt"
	"net/http"
	"time"

	"diploma/internal/pkg/constants"
	"diploma/internal/planner"
)

func HandlePlanner(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		query := req.URL.Query()

		now, err := time.Parse(constants.DateFormat, query.Get("now"))
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

		response, err := planner.NextDate(now, dtstart, repeat)

		if err != nil {
			writeError(res, fmt.Sprint(err), 400)
			return
		}

		writeResponse(res, response, http.StatusCreated)
		return
	}

	writeResponse(res, nil, http.StatusMethodNotAllowed)
}

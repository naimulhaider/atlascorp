package api

import (
	"encoding/json"
	"net/http"

	"github.com/naimulhaider/atlascorp/actions"
	"github.com/naimulhaider/atlascorp/config"
)

func HandleDNS(w http.ResponseWriter, r *http.Request) {
	rData := DNSRequest{}
	err := json.NewDecoder(r.Body).Decode(&rData)
	if err != nil {
		ServeJSON(w, http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	dm, err := rData.ToModel()
	if err != nil {
		ServeJSON(w, http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	loc := actions.CalculateLocation(config.SectorID, dm.X, dm.Y, dm.Z, dm.Velocity)

	res := DNSResponse{
		Loc: loc,
	}

	ServeJSON(w, http.StatusOK, res)
}

func ServeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.Write(b)
}

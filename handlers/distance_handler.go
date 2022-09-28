package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DelusionalOptimist/distance-api/utils"
	"googlemaps.github.io/maps"
)

type DistanceResponse struct {
}

func (h *Handler) GetDistance (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		h.Log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	distanceMatrixReq := &maps.DistanceMatrixRequest{
		Origins: []string{
			r.FormValue("origin"),
		},
		Destinations: []string{
			r.FormValue("destination"),
		},
	}

	distanceMatrixResp, err := utils.GoogleDistanceAPI(distanceMatrixReq)
	if err != nil {
		h.Log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(distanceMatrixResp)
	if err != nil {
		h.Log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

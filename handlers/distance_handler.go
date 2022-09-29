package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DelusionalOptimist/distance-api/utils"
	"googlemaps.github.io/maps"
)

type DistanceRequest struct {
	Origin string `json:"origin"`
	Destination string `json:"destination"`
}

type DistanceResponse struct {
	Status string `json:"status"`
	Duration string `json:"duration"`
	Distance string `json:"distance"`
}

func (h *Handler) GetDistance (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	distanceReqeust := &DistanceRequest{}
	err := json.NewDecoder(r.Body).Decode(distanceReqeust)
	if err != nil {
		h.Log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	distanceMatrixReq := &maps.DistanceMatrixRequest{
		Origins: []string{
			distanceReqeust.Origin,
		},
		Destinations: []string{
			distanceReqeust.Destination,
		},
	}

	googleAPIResponse, err := utils.GoogleDistanceAPI(distanceMatrixReq)
	if err != nil {
		h.Log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	distanceMatrixResp := &DistanceResponse{}
	if googleAPIResponse.Rows[0].Elements[0].Status == "OK" {
		distanceMatrixResp = &DistanceResponse{
			Status: "OK",
			Distance: googleAPIResponse.Rows[0].Elements[0].Distance.HumanReadable,
			Duration: googleAPIResponse.Rows[0].Elements[0].Duration.String(),
		}
	} else {
		distanceMatrixResp = &DistanceResponse{
			Status: "ZERO_RESULTS",
			Distance: "undefined",
			Duration: "undefined",
		}
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

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DelusionalOptimist/distance-api/utils"
	"googlemaps.github.io/maps"
)

// the JSON received from frontend
type DistanceRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// object containing distance sent to frontend
type DistanceResponse struct {
	// Status:
	// OK- a route was found between the origin and destination
	// ZERO_RESULTS: no route found between origin and destination
	Status string `json:"status"`

	// distance in km
	Distance string `json:"distance"`

	// time taken to reach by road
	// in human readable format
	Duration string `json:"duration"`
}

// returns DistanceResponse
func (h *Handler) GetDistance(w http.ResponseWriter, r *http.Request) {
	// only POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// read the incoming request body
	distanceReqeust := &DistanceRequest{}
	err := json.NewDecoder(r.Body).Decode(distanceReqeust)
	if err != nil {
		h.Log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create a request to the GoogleAPI
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

	// validate result
	distanceMatrixResp := &DistanceResponse{}
	if googleAPIResponse.Rows[0].Elements[0].Status == "OK" {
		distanceMatrixResp = &DistanceResponse{
			Status:   "OK",
			Distance: googleAPIResponse.Rows[0].Elements[0].Distance.HumanReadable,
			Duration: googleAPIResponse.Rows[0].Elements[0].Duration.String(),
		}
	} else {
		distanceMatrixResp = &DistanceResponse{
			Status:   "ZERO_RESULTS",
			Distance: "undefined",
			Duration: "undefined",
		}
	}

	// write back JSON response containing duration and distance
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

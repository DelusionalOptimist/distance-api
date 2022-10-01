package utils

import (
	"context"

	"googlemaps.github.io/maps"
)

// creates a call to the GoogleDistanceAPI
// returns the response to calling function as it is
// return value may change in future
func GoogleDistanceAPI(distanceRequest *maps.DistanceMatrixRequest) (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey(GlobalConfig.APIKey))
	if err != nil {
		return nil, err
	}

	distanceResponse, err := c.DistanceMatrix(context.Background(), distanceRequest)
	if err != nil {
		return nil, err
	}

	return distanceResponse, nil
}

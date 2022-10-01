package utils

import (
	"context"

	"googlemaps.github.io/maps"
)

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

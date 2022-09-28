package utils

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

func GoogleDistanceAPI(distanceRequest *maps.DistanceMatrixRequest) (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey("<api-key>"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil, err
	}

	distanceResponse, err := c.DistanceMatrix(context.Background(), distanceRequest)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return nil, err
	}

	//pretty.Println(distance)
	return distanceResponse, nil
}

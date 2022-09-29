include .api-key

PORT="8080"

run:
	go build -o distance-api .
	PORT=$(PORT) API_KEY=$(API_KEY) ./distance-api

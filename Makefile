include .api-key

PORT="8080"

run:
	go build -o distance-api .
	PORT=$(PORT) API_KEY=$(API_KEY) ./distance-api

docker-build:
	docker build -t distance-api:latest --build-arg API_KEY=$(API_KEY) --build-arg PORT=$(PORT) .

docker-build-prod:
	docker build -t delusionaloptimist/distance-api:latest --build-arg API_KEY=$(API_KEY) --build-arg PORT=$(PORT) .
	docker push delusionaloptimist/distance-api:latest

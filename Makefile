build:
	go build -o walkara ./cmd/api

run:
	go run ./cmd/api

docker-build:
	docker build -t walkara .

docker-run:
	docker run -p 8080:8080 walkara
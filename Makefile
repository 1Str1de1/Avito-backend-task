.PHONY: build, test
build:
	go build -v ./cmd/server
docker-restart:
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d
DEFAULT: build

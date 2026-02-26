APP_NAME=desent-solutions
BINARY_NAME=app
MAIN_PATH=./cmd/server
PORT=8080
SPEEDRUN_SCRIPT=./speedrun.sh

.PHONY: help build run test clean docker-build docker-run

help:
	@echo "Available commands:"
	@echo "  make build         - Build binary"
	@echo "  make run           - Run locally"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean binary"
	@echo "  make docker-build  - Build docker image"
	@echo "  make docker-run    - Run docker container"

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	rm -f $(BINARY_NAME)

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p $(PORT):8080 $(APP_NAME)

speedrun:
	@test -f $(SPEEDRUN_SCRIPT) || (echo "speedrun.sh not found"; exit 1)
	chmod +x $(SPEEDRUN_SCRIPT)
	$(SPEEDRUN_SCRIPT)
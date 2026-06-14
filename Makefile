APP_NAME=tasks
BIN_DIR=bin

.PHONY: build run test install clean deps

deps:
	go mod tidy
	go mod download

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) .

run:
	go run .

test:
	go test ./...

install:
	go install .

clean:
	rm -rf $(BIN_DIR)
	rm -f tasks
	rm -f *.exe

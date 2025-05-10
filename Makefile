APP_NAME := bank-account-system
MAIN_FILE := main.go
BUILD_DIR := build

.PHONY: all fmt test build run clean swagger

all: build

fmt:
	go fmt ./...

test:
	go test -v ./...

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

run: build
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	go clean
	rm -rf $(BUILD_DIR)

swagger:
	swagger generate spec -o ./docs/swagger.json --scan-models

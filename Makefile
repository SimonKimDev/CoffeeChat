APP_NAME := coffeechat 
MAIN_PATH := ./cmd/$(APP_NAME)
BIN_PATH := ./bin/$(APP_NAME)
PGK := ./...
PORT := 8080


build:
	go build -o $(BIN_PATH) $(MAIN_PATH)

run: build
	./$(BIN_PATH)

test:
	go test -v ./...

clean:
	rm -f $(BIN_PATH) 

.PHONY: build run clean test

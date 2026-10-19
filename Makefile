APP_NAME  := coffeechat-server
MAIN_PATH := ./cmd/server
BIN_PATH  := ./bin/$(APP_NAME)

build:
	go build -o $(BIN_PATH) $(MAIN_PATH)

run: build
	$(BIN_PATH)

test:
	go test ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf ./bin

.PHONY: build run test vet tidy clean

BIN     := maximus
MODULE  := github.com/maximus
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

.PHONY: build test lint clean docker-up docker-down

build:
	go build $(LDFLAGS) -o $(BIN) ./cmd/maximus

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -f $(BIN)

docker-up:
	docker compose up -d

docker-down:
	docker compose down -v

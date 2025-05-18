BINARY   := gshortcuts
CMD_PATH := ./cmd/gshortcuts

.PHONY: all build run clean

all: build

build:
	go vet ./...
	go build -o bin/$(BINARY) $(CMD_PATH)

run:
	go run $(CMD_PATH)

install: build
	go install $(CMD_PATH)

clean:
	rm -rf bin/


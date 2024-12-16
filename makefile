.PHONY: all run build test clean

all: build

run: build
	./llp

build:
	go build --tags "fts5" -o llp main.go

test:
	go test --tags "fts5" -v ./...

clean:
	go clean
	rm -f llp

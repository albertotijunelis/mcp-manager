.PHONY: build install test lint clean release-dry run

build:
	go build -o bin/mcp .

install:
	go install .

test:
	go test ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/ dist/

release-dry:
	goreleaser release --snapshot --clean

run:
	go run main.go

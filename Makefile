build:
	@go build -o bin/goREST ./cmd

run: build
	@./bin/goREST

test:
	@go test -v ./...
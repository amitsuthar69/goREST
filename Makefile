build:
	@go build -o bin/goREST

run: build
	@./bin/goREST

test:
	@go test -v ./...
build:
	@go build -o ./bin/go-api.exe ./cmd/go-api/main.go

run: build
	@./bin/go-api.exe
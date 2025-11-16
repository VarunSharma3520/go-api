build:
	@go build -o ./bin/app.exe ./cmd/go-api/main.go

run: build
	@./bin/app.exe
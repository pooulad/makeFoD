build:
	@go build -o ./bin/makeFoD ./cmd/app/main.go

run: build
	@./bin/makeFoD

tidy:
	@go mod tidy

testgo:
	@go test ./... -v
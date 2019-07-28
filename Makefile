run-api:
	@go run ./cmd/api.go

run-server:
	@go run ./internal/server/main.go

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .
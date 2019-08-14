run-api:
	@go run ./cmd/api.go

run-server:
	@go run ./internal/server/main.go

api-integration-tests:
	@go test ./internal/api/integration-tests/...

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .
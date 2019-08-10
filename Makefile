run-api:
	@go run ./cmd/api.go

run-server:
	@go run ./internal/server/main.go

integration-tests:
	@go test do/internal/api/integration-tests

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .
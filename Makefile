run-api:
	@go run ./cmd/api/main.go

api-integration-tests:
	@go test -v ./internal/api/integration-tests/...

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .


run-server:
	@go run ./cmd/server/main.go

build-server-container:
	@docker build -f ./build/package/server/Dockerfile -t beegone-server .
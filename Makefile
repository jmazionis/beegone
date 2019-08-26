run-api:
	@go run ./cmd/api/main.go

api-unit-tests: 
	@go test

api-integration-tests:
	@go test -v ./internal/api/integration-tests/...

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .

build-server:
	@go build -o ./bin/server ./cmd/server/...
	@./bin/server

build-server-container:
	@docker build -f ./build/package/server/Dockerfile -t beegone-server .

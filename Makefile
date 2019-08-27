run-api:
	@go run ./cmd/api/main.go

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .

run-server:
	@go build -o ./bin/server ./cmd/server/...
	@./bin/server

build-server-container:
	@docker build -f ./build/package/server/Dockerfile -t beegone-server .

run-unit-tests: 
	@go vet ./...
	@go test -v -tags=unit ./...

run-integration-tests:
	@go test -v -tags=integration ./...
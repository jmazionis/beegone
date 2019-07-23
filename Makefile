run-api:
	@go run ./cmd/api.go

build-api-container:
	@docker build -f ./build/package/api/Dockerfile -t beegone-api .
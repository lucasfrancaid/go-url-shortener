port ?= 3333

http:
	@echo "--> Starting HTTP Stantard Server"
	@export PORT=$(port) \
			DOMAIN=$(domain) \
			REPOSITORY_ADAPTER=$(repository) \
			&& go run cmd/http/standard.go

test:
	@echo "--> Running Tests"
	@go test ./... $(args)

coverage:
	@echo "--> Running Tests and Coverage"
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
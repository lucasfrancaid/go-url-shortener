env ?= dev
port ?= 3333

http:
	@echo "--> Starting HTTP Stantard Server"
	@ENV=$(env) go run cmd/http/standard.go

test:
	@echo "--> Running Tests"
	@ENV=test go test ./... $(args)

coverage:
	@echo "--> Running Tests and Coverage"
	@ENV=test go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
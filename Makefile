env ?= dev
port ?= 3333
adapter ?= standard

http:
	@echo "--> Starting HTTP $(adapter) server"
	@ENV=$(env) go run cmd/http/$(adapter)/main.go

test:
	@echo "--> Running Tests"
	@ENV=test go test ./... $(args)

coverage:
	@echo "--> Running Tests and Coverage"
	@ENV=test go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
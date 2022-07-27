## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool Speedlight and copies it to myapp
build_cli:
	@go mod vendor
	@go build -o ../app/speedlight ./cmd/cli

## build: builds the command line tool to dist directory
build:
	@go mod vendor
	@go build -o ./dist/speedlight ./cmd/cli
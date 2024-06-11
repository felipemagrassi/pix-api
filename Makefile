fmt:
	go fmt ./...

test: go test ./... --coverprofile coverage.out

up:
	docker-compose up -d --build

cover:
	go tool cover -html coverage.out

install-dependencies:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag --version

install-all: install-binaries

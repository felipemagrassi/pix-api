.PHONY: fmt test up seed cover install-dependencies install-all

fmt:
	go fmt ./...

test: 
	go test ./... --coverprofile coverage.out

up:
	docker-compose up -d --build

seed: 
	docker-compose exec app go run scripts/seed.go

cover:
	go tool cover -html coverage.out

install-dependencies:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag --version

install-all: install-binaries

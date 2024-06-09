fmt:
	go fmt ./...

mock:
	go generate -v ./...

test: mock
	go test ./... --coverprofile coverage.out

cover:
	go tool cover -html coverage.out

install-dependencies:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/mock/mockgen@latest
	swag --version
	mockgen --version

install-all: install-binaries

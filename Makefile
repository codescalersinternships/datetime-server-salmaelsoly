build :
	go build -o ./builds/gin ./cmd/gin/main.go
	go build -o ./builds/http ./cmd/http/main.go
format:
	go fmt ./...
lints:
	golangci-lint run
run-tests:
	go test ./... -v
build-docker:
	docker compose up
all: build format lints run-tests build-docker 
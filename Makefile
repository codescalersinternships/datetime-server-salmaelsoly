build :
	go build -o ./builds/gin ./cmd/gin/main.go
	go build -o ./builds/http ./cmd/http/main.go
format:
	go fmt ./...
lints:
	golangci-lint run
run-tests:
	go test ./... -v
build-image-docker:
	docker build -f Dockerfile.http -t http-server .
	docker build -f Dockerfile.gin -t gin-server .

run-image-docker:
	docker run -d -p 3000:8080 http-server
	docker run -d -p 5000:8080 gin-server
all: build format lints run-tests build-image-docker run-image-docker
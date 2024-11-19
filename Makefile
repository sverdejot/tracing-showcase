fmt:
	@go fmt ./...

vet: fmt
	@go vet ./...

build: vet
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/server ./cmd/server/main.go && \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/client ./cmd/client/main.go

run: build
	@docker compose up

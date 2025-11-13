
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/mobile-numbers-mine-linux-amd64 cmd/main.go

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/mobile-numbers-mine-darwin-amd64 cmd/main.go

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o bin/mobile-numbers-mine-windows-amd64.exe cmd/main.go
	
init-run:
	@cp config.yaml-sample config.yaml
	@echo "Starting the server..."
	go run cmd/main.go

run:
	@echo "Starting the server..."
	REDIS_ADDR=localhost:6379 REDIS_DB=0 REDIS_PASS= go run cmd/main.go

doc:
	swag init -g cmd/main.go -o docs/api	

build-all: build-linux-amd64 build-darwin-amd64 build-windows-amd64

clean:
	@rm -rf bin/*
	@rm config.yaml

.PHONY: build-linux-amd64 build-darwin-amd64 build-windows-amd64 run init-run doc build-all clean

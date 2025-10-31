go_bin_dir = "/Users/ashva/Documents/go_projects/bin/"

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
	go run cmd/main.go

doc:
	${go_bin_dir}swag init -g cmd/main.go -o docs/api	

build-all: build-linux-amd64 build-darwin-amd64 build-windows-amd64

clean:
	@rm -rf bin/*
	@rm config.yaml

.PHONY: build-linux-amd64 build-darwin-amd64 build-windows-amd64 run init-run doc build-all clean

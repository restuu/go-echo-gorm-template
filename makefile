.PHONY: build-windows generate lint

build-windows:
	go build -o bin/webserver.exe .\cmd\webserver\...

generate:
	go generate ./...

lint:
	golangci-lint run ./... --fix
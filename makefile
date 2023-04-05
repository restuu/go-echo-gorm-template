.PHONY: build-windows

build-windows:
	go build -o bin/webserver.exe .\cmd\webserver\...
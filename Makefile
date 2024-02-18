.PHONY: build-windows build-linux clean

# Go source files
SRC := ./cmd/ratTask/main.go

# Binary name
BINARY := rattask

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY)_windows_amd64.exe $(SRC)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY)_linux_amd64 $(SRC)

clean:
	rm -f ./bin/$(BINARY)_windows_amd64.exe ./bin/$(BINARY)_linux_amd64

build: build-windows build-linux

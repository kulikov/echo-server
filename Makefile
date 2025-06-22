TARGET_OS=linux
TARGET_ARCH=amd64

export CGO_ENABLED=0

build:
	GOOS=linux GOARCH=amd64 go build -v -trimpath -ldflags "-s -w" -o ./bin/echo-server main.go

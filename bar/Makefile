all: clean build

build:
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o docker/main cmd/api/main.go
	
clean:
	rm -f docker/main

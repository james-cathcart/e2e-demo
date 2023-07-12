all: clean build

build:
	env GOOS=linux go build -o docker/main cmd/api/main.go

clean:
	rm -f docker/main

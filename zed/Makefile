all: clean build

build:
	env GOOS=linux go build -o docker/main cmd/api/main.go
	cp ads.db docker/ads.db

clean:
	rm -f docker/main
	rm -f docker/ads.db

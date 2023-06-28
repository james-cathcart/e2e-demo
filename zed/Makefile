all: clean build

build:
	GOOS=linux GOARCH=amd64 go build -o docker/amd64/main cmd/api/main.go
	GOOS=darwin GOARCH=arm64 go build -o docker/arm64/main cmd/api/main.go
	cp ads.db docker/amd64/ads.db
	cp ads.db docker/arm64/ads.db

clean:
	rm -f docker/amd64/main
	rm -f docker/amd64/ads.db
	rm -f docker/arm64/main
	rm -f docker/arm64/ads.db
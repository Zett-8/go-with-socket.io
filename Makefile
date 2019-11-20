start:
	docker-compose up

build-for-mac:
	env GOOS=darwin GOARCH=amd64 go build -o app-mac .

build-for-linux:
	env GOOS=linux GOARCH=amd64 go build -o app-linux .
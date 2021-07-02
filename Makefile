.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run: build
	docker-compose up -d --remove-orphans --build server 

test:
	go test ./...

down:
	docker stop server && docker rm server
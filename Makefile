.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run: build
	docker-compose -f dev-docker-compose.yaml up -d

test:
	go test ./...

stop:
	docker stop server && docker rm server

rebuild: stop
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go && docker-compose up -d --remove-orphans --build server

migrate-up:
	 migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" up

migrate-drop:
	 migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" drop

migrate-down:
	 migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" down

.DEFAULT_GOAL := run
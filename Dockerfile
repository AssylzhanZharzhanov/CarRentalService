FROM golang:alpine

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

RUN go mod download

ENTRYPOINT CompileDaemon --build="go build -o ./.bin/app ./cmd/api/main.go" --command=./.bin/app
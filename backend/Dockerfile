FROM golang:1.14-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir /app
WORKDIR /app

COPY /backend/go.mod /backend/go.sum ./

RUN go mod download

COPY ./backend .

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o main.exe" --command=./main.exe
FROM golang:latest

WORKDIR /app

COPY ./ /app

ENV DB_HOST
ENV DB_USER
ENV DB_PW
ENV DB_NAME
ENV DB_PORT

RUN go mod download

ENTRYPOINT go run main.go

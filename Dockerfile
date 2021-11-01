FROM golang:latest

WORKDIR /app

COPY ./ /app

ENV DB_HOST=34.89.238.138
ENV DB_USER=myuser
ENV DB_PW=22papich87
ENV DB_NAME=postgres
ENV DB_PORT=5432

RUN go mod download

EXPOSE 8181

ENTRYPOINT go run main.go

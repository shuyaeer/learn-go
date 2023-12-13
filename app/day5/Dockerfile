FROM golang:latest as build

WORKDIR /app

COPY go.mod go.sum ./

COPY main.go .
RUN GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o main main.go

ENTRYPOINT [ "./main" ]

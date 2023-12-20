FROM golang:1.21.5
WORKDIR /app
RUN go mod init echo_app

RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware

COPY . .
CMD ["go", "run", "main.go"]

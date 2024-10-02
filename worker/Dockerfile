FROM golang:1.23.2

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build cmd/tiksup/main.go

ENTRYPOINT [ "./main" ]

FROM golang:1.23.2-alpine3.20 AS build

RUN apk add --no-cache librdkafka-dev gcc libc-dev

WORKDIR /app

COPY . .

ENV CGO_ENABLED=1

RUN go mod tidy

RUN go build -tags musl cmd/tiksup/main.go

FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY --from=build /app/main .

ENTRYPOINT [ "./main" ]

FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

ENV CGO_ENABLED 0

ENV GOOS linux

RUN go mod download

COPY . .

RUN ls

RUN go build -o tiny src/cmd/server/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/tiny .
COPY --from=builder /app/src/cmd/server/config ./config
RUN ls
EXPOSE 8080

CMD ["./tiny"]


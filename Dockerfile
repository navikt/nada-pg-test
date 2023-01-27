FROM golang:1.19-alpine as builder

ENV CGO_ENABLED=0

WORKDIR /src
COPY go.sum go.sum
COPY go.mod go.mod
RUN go mod download
COPY . .
RUN go build -o nada-db-test .

FROM golang:1.19-alpine
LABEL org.opencontainers.image.source https://github.com/navikt/nada-pg-test

WORKDIR /app
COPY --from=builder /src/nada-db-test /app/nada-db-test
CMD ["/app/nada-db-test"]
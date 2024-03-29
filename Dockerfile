FROM golang:1.21-alpine as builder

WORKDIR /src
COPY go.sum go.sum
COPY go.mod go.mod

RUN go mod download

COPY main.go main.go
COPY pkg pkg

RUN CGO_ENABLED=0 go build -o app .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /src/app /app/app
CMD ["/app"]

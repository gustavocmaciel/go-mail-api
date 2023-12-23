# Build stage
FROM golang:latest AS builder

WORKDIR /app

COPY . .

COPY go.mod go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-mail-api ./cmd/main.go

# Final stage
FROM scratch

WORKDIR /

COPY --from=builder /go-mail-api .

CMD ["/go-mail-api"]
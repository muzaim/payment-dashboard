FROM golang:1.25-alpine AS builder

WORKDIR /app

# Required for CGO + sqlite
RUN apk add --no-cache gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# IMPORTANT: CGO must be ENABLED
RUN CGO_ENABLED=1 go build -o app


FROM alpine:latest

WORKDIR /app

# sqlite runtime + certs
RUN apk add --no-cache ca-certificates sqlite-libs

COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]

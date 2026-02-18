FROM golang:1.25.4-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o out/server ./cmd

FROM scratch
WORKDIR /app
COPY --from=builder /app/out/server /app/server
COPY --from=builder /app/web /app/web
ENTRYPOINT ["/app/server"]

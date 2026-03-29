FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY relay.go .
RUN go mod init relay && go get github.com/hashicorp/yamux
RUN CGO_ENABLED=0 GOOS=linux go build -o relay-server .

FROM scratch
COPY --from=builder /app/relay-server /relay-server
EXPOSE 7000 8000
ENTRYPOINT ["/relay-server"]
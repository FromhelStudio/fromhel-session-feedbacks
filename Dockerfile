FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fromhel-session-feedbacks ./cmd/server/main.go

FROM ubuntu:latest  

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 8080
WORKDIR /app
COPY --from=builder /app/fromhel-session-feedbacks .
COPY --from=builder /app/.env .

CMD ["/app/fromhel-session-feedbacks"]

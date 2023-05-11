FROM golang:1.19 AS builder
WORKDIR /app
ENV GOPATH=/
COPY . .
RUN apt-get update &&\
 go mod download &&\
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o todo-app ./cmd/main.go


FROM alpine:latest as production
WORKDIR /app
RUN apk add --no-cache postgresql-client
COPY config.yaml wait-postgres.sh ./
RUN chmod +x wait-postgres.sh
COPY --from=builder /app/todo-app .
CMD ["./todo-app"]

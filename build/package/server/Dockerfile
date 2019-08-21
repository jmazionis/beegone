FROM golang:1.12 AS gobuilder
WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY ./internal/server ./internal/server
COPY ./cmd/server/main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin main.go

CMD ["./bin"]

EXPOSE 9090
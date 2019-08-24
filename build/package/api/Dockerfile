FROM golang:1.12 AS gobuilder
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./internal/api ./internal/api
COPY ./cmd/api/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin main.go

FROM alpine
WORKDIR /app
COPY --from=gobuilder /app/bin /app/bin
CMD ["./bin"]
EXPOSE 9090

FROM node:10.16-alpine AS nodebuilder
WORKDIR /app
COPY ./web/app/package.json ./web/app/yarn.lock ./web/app/tsconfig.json ./
RUN yarn install
COPY ./web/app/src ./src
COPY ./web/app/public ./public
RUN yarn build

FROM golang:1.12 AS gobuilder
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./internal/server ./internal/server
COPY ./cmd/server/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/server main.go

FROM alpine
WORKDIR /app
COPY --from=gobuilder /app/bin /app/bin
COPY --from=nodebuilder /app/build /app/web/app/build
CMD ["./bin/server"]
EXPOSE 9000
# Dev
FROM golang:1.23-alpine AS dev

WORKDIR /var/www/mindlab

RUN apk add --no-cache curl bash ca-certificates git openssh make gcc musl-dev

CMD ["go", "run", "./cmd/app/main.go"]

# Build ---
FROM golang:1.23-alpine AS build

WORKDIR /src

RUN apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/app ./cmd/app/main.go

# Prod
FROM alpine:latest AS runtime

WORKDIR /app

RUN addgroup -g 1000 go-user && \
    adduser -u 1000 -G go-user -s /bin/sh -D go-user

COPY --from=build /app/app ./app
COPY --from=build /src/build/.env ./cmd/app/.env

USER go-user

CMD ["./app"]
FROM golang:1.23-alpine AS build

WORKDIR /var/www/mindlab

RUN apk update && \
    apk add --no-cache curl bash ca-certificates git openssh make gcc musl-dev && \
    rm -fr /var/cache/apk/*

COPY . .
RUN go mod download


RUN addgroup -g 1000 go-user \
    && adduser -u 1000 -G go-user -s /bin/sh -D go-user



RUN go build -o /var/www/mindlab/cmd/app/ -v ./cmd/app/main.go


RUN chown -R go-user ./

FROM build AS base

WORKDIR /app


COPY --from=build /var/www/mindlab/cmd/app/ ./cmd/app

COPY --from=build /var/www/mindlab/build/.env ./cmd/app/.env


RUN chown -R go-user ./

CMD ["tail", "-f", "/dev/null"]

USER go-user


FROM base AS app


WORKDIR /app/cmd/app/
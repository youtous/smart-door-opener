FROM golang:1.21-alpine AS builder

WORKDIR /src/app

COPY go.mod go.sum ./
RUN apk add --no-cache git \
    && go mod download

COPY . .
RUN go install

FROM alpine:latest
LABEL maintainer "youtous <contact@youtous.me>"

EXPOSE 1323

WORKDIR "/app"

ENTRYPOINT ["/app/smart-door-opener", "server"]

RUN apk add --no-cache ca-certificates curl

HEALTHCHECK CMD curl --silent --fail http://localhost:1323/_health > /dev/null || exit 1

COPY --from=builder /go/bin/smart-door-opener /app/
COPY static /app/static
COPY views /app/views

USER 1000
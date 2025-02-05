# syntax=docker/dockerfile:1

FROM golang:1.23.3-alpine AS builder

ARG APP_VERSION="undefined"
ARG BUILD_TIME="undefined"

WORKDIR /go/src/github.com/ci-space/version-badge

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -ldflags="-s -w -X 'main.Version=${APP_VERSION}' -X 'main.BuildDate=${BUILD_TIME}'" -o /go/bin/version-badge /go/src/github.com/ci-space/version-badge/main.go

######################################################

FROM alpine

RUN apk add tini

COPY --from=builder /go/bin/version-badge /go/bin/version-badge

# https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.title="version-badge"
LABEL org.opencontainers.image.description="is console app and GitHub action for generate SVG badges with version of language or dependency"
LABEL org.opencontainers.image.url="https://github.com/ci-space/version-object"
LABEL org.opencontainers.image.source="https://github.com/ci-space/version-object"
LABEL org.opencontainers.image.vendor="ArtARTs36"
LABEL org.opencontainers.image.version="${APP_VERSION}"
LABEL org.opencontainers.image.created="${BUILD_TIME}"
LABEL org.opencontainers.image.licenses="MIT"

COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x ./docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

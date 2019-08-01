# stage 1: build
FROM golang:1.10-alpine AS builder
LABEL maintainer="nightfury1204"

# Add source code
RUN mkdir -p /go/src/github.com/nightfury1204/movie-listing-app
ADD . /go/src/github.com/nightfury1204/movie-listing-app

# Build binary
RUN cd /go/src/github.com/nightfury1204/movie-listing-app && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/movie-listing-app

# stage 2: lightweight "release"
FROM alpine:latest
LABEL maintainer="nightfury1204"

COPY --from=builder /go/bin/movie-listing-app /app/
COPY --from=builder /go/src/github.com/nightfury1204/movie-listing-app/templates /app/templates

EXPOSE 8443

WORKDIR /app

ENTRYPOINT [ "./movie-listing-app", "run"]

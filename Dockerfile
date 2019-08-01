# stage 1: build
FROM golang:1.10-alpine AS builder
LABEL maintainer="nightfury1204"

# Add source code
RUN mkdir -p /go/src/github.com/nightfury1204/movie-search-app
ADD . /go/src/github.com/nightfury1204/movie-search-app

# Build binary
RUN cd /go/src/github.com/nightfury1204/movie-search-app && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/movie-search-app

# stage 2: lightweight "release"
FROM alpine:latest
LABEL maintainer="nightfury1204"

COPY --from=builder /go/bin/movie-search-app /app/
COPY --from=builder /go/src/github.com/nightfury1204/movie-search-app/templates /app/templates

EXPOSE 8443

WORKDIR /app

ENTRYPOINT [ "./movie-search-app", "run"]

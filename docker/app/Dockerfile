# Start from golang base image
FROM golang:alpine AS builder

ARG APP_PORT

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Reza Fahlevi <rzfhlv.me@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git gcc musl-dev

# Set the current working directory inside the container 
WORKDIR /gin-example

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN go build -ldflags ' -w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Start a new stage from scratch
FROM alpine:latest

RUN apk update && apk add --no-cache bash mysql-client

# RUN adduser -D gin
# USER gin

WORKDIR /gin-example/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /gin-example/bin/app /gin-example/
COPY --from=builder /gin-example/bin/migrate /gin-example/
COPY --from=builder /gin-example/database/migrations /gin-example/database/migrations
COPY --from=builder /gin-example/.env /gin-example/

COPY --from=builder /gin-example/docker/app/bin /usr/local/bin/gin-example/
RUN chmod +x /usr/local/bin/gin-example/*

# Expose port to the outside world
EXPOSE $APP_PORT

#Command to run the executable
CMD ["/gin-example/app"]
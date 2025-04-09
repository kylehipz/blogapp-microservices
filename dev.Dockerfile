FROM golang:1.24-alpine

WORKDIR /app

# Install Air
RUN go install github.com/air-verse/air@latest

ARG DIR

# Copy go.work and go.mod files
COPY go.work go.work.sum ./
COPY libs/go.mod libs/
COPY auth/go.mod auth/
COPY blogs/go.mod blogs/
COPY follow/go.mod follow/
COPY home-feed/go.mod home-feed/
COPY search/go.mod search/

# Download dependencies
RUN go mod download

# Copy entire source code
COPY libs libs
COPY $DIR $DIR

WORKDIR /app/$DIR

CMD ["air"]


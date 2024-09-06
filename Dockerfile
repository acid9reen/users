FROM golang:1.22.6-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.22.6-alpine AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY ./cmd ./cmd
COPY ./config ./config
COPY ./go.mod .
COPY ./go.sum .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main cmd/httpserver/main.go

FROM debian:bookworm-slim

WORKDIR /app

ARG UID=10000
ARG GID=10001
ARG USERNAME=app
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        sudo \
        curl \
    && addgroup --gid $GID ${USERNAME} && \
    adduser --uid $UID --gid $GID ${USERNAME}

# Copy the built application
COPY ./static ./static
COPY --from=builder /app/main .

# Change ownership of the application binary
RUN chown ${UID}:${GID} /app

USER ${USERNAME}
CMD ["./main"]

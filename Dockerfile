FROM golang:alpine AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /build

# Download go dependencies
COPY go.mod .
RUN go mod download

# Copy into the container
COPY . .

# Statically build the application
RUN go build -o zippy2 -trimpath -ldflags "-s -w -extldflags '-static'" .

# Build final image using nothing but the binary
FROM alpine:3.17.2

LABEL org.opencontainers.image.source=https://github.com/kubefirst-demo-bot/zippy2
LABEL org.opencontainers.image.description="simple golang http server"

COPY --from=builder /build/zippy2 /usr/bin/

ENTRYPOINT ["/usr/bin/zippy2"]

## Containerfile
# ----------------------------------------------
# Build Stage
#
FROM golang:latest AS builder

WORKDIR /workspaces/websb

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make test build

# ----------------------------------------------
# Package Stage
FROM debian:stable-slim

WORKDIR /app
COPY --from=builder /workspaces/websb/dist/websb .
COPY --from=builder /workspaces/websb/etc ./etc
VOLUME [ "/etc" ]
EXPOSE 8080/tcp
EXPOSE 8443/tcp
CMD [ "/app/websb", "daemon" ]

# Alpine is chosen for its small footprint compared to Ubuntu
FROM golang:rc-alpine as builder
RUN	apk add --no-cache ca-certificates

WORKDIR /app

# Download necessary Go modules
COPY go.mod /app/
COPY go.sum /app/
RUN go mod download

# Copy serv-sub source code
COPY ./cmd/serv-sub/*.go /app/

# Build
RUN set -x && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /serv-sub -buildvcs=false

# Copy files to final location
FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /serv-sub /usr/bin/serv-sub
ENTRYPOINT [ "serv-sub" ]
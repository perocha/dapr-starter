# Alpine is chosen for its small footprint compared to Ubuntu

# Step #1 - Download all necessary GO modules
FROM golang:rc-alpine as modules
WORKDIR /modules
COPY go.mod go.sum /modules/
RUN go mod download

# Step #2 - Build the application
FROM golang:rc-alpine as builder
RUN	apk add --no-cache ca-certificates
COPY --from=modules /go/pkg /go/pkg
COPY . /app
COPY ./cmd/serv-sub/*.go /app/
WORKDIR /app
RUN set -x && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app -buildvcs=false
#COPY ./config/* /config/

# Build
#RUN set -x && env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /serv-sub -buildvcs=false

# Copy files to final location
FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
ENTRYPOINT [ "serv-sub" ]
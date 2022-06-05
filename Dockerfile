# Step 1: Modules caching
FROM golang:rc-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:rc-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
COPY ./cmd/serv-sub/*.go /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o serv-sub -buildvcs=false

# Step 3: Final
FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /app/serv-sub /usr/bin/serv-sub
ENTRYPOINT ["serv-sub"]
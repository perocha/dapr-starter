# Step 1: Modules caching
FROM golang:rc-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:rc-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /serv-sub
WORKDIR /serv-sub
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/serv-sub ./cmd/serv-sub

# Step 3: Final
FROM scratch
COPY --from=builder /serv-sub/config /config
COPY --from=builder /bin/serv-sub /serv-sub
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
CMD ["/serv-sub"]
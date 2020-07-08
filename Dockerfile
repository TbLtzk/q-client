# Build Geth in a stock Go builder container
FROM golang:1.14.4-alpine3.12 as builder

ADD . /go-ethereum

# Pull Geth into a second stage deploy alpine container
FROM alpine:3.12.0

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 8547 30303 30303/udp
ENTRYPOINT ["geth"]

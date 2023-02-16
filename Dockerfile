# Support setting various labels on the final image
ARG ARCH=""
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.19-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

ADD . /q-client
ARG BUILD_TOKEN
ARG USERNAME=oauth2
RUN git config --global url."https://${USERNAME}:${BUILD_TOKEN}@gitlab.com/".insteadOf https://gitlab.com/
RUN go env -w GOPRIVATE=gitlab.com/q-dev/*
RUN cd /q-client && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /q-client/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth"]

# Add some metadata labels to help programatic image consumption
ARG ARCH=""
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"

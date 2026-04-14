# Support setting various labels on the final image
ARG ARCH=
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

# Build Geth in a stock Go builder container
FROM golang:1.21-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

# Copy module files first so we can download deps (and so verify has go.mod)
COPY go.mod go.sum ./

# Verify and download dependencies (use module cache; vendor prunes C sources for blst - golang/go#26366)
RUN go mod verify && go mod download

# Build geth (use -mod=mod so module cache is used; vendor does not include blst C headers)
ADD . /q-client
ARG BUILD_TOKEN
ARG USERNAME=oauth2
RUN git config --global url."https://${USERNAME}:${BUILD_TOKEN}@gitlab.com/".insteadOf https://gitlab.com/
RUN go env -w GOPRIVATE=gitlab.com/q-dev/*
RUN cd /q-client && GOFLAGS=-mod=mod make geth

# Pull Geth into a second stage deploy alpine container
FROM ${ARCH}alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /q-client/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
ENTRYPOINT ["geth"]

# Add some metadata labels to help programatic image consumption
ARG ARCH=
ARG COMMIT=""
ARG VERSION=""
ARG BUILDNUM=""

LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"

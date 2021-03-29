FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/cosmos/cosmos-sdk

# Add source files
COPY . .

RUN make build-simd

# Final image
FROM alpine:edge
LABEL org.opencontainers.image.source https://github.com/adoriasoft/polkadot_cosmos_integration

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/nsd /usr/bin/nsd
COPY --from=build-env /go/bin/nscli /usr/bin/nscli
COPY --from=build-env /go/src/github.com/cosmos/cosmos-sdk/simapp/init.sh /usr/bin/init.sh

RUN /bin/sh /usr/bin/init.sh

EXPOSE 26656 26657 26658 1317 9090

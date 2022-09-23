FROM golang:1.17 as builder

# create a working directory
COPY . /photon
WORKDIR /photon


# Build binary
RUN GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o ph
# Uncomment only when build is highly stable. Compress binary.
# RUN strip --strip-unneeded ts
# RUN upx ts

# use a minimal alpine image
FROM alpine:3.12
# add ca-certificates in case you need them
RUN apk update && apk add ca-certificates libc6-compat && rm -rf /var/cache/apk/*
# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /photon/ph .
# run the binary
CMD ["./ph"]
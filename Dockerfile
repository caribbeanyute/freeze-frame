FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git ffmpeg
WORKDIR $GOPATH/src/freezeframe/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/freezeframe
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/freezeframe /go/bin/freezeframe
EXPOSE 3001
# Run the hello binary.
ENTRYPOINT ["/go/bin/freezeframe"]
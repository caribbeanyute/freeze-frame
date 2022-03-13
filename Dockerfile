FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .

# Fetch dependencies.
RUN go mod download


# Build the binary.
RUN CGO_ENABLED=0 go build -o freezeframe
############################
# STEP 2 build a small image
############################
FROM jrottenberg/ffmpeg:4.4-scratch
# Copy our static executable.
COPY --from=builder /app/freezeframe /
EXPOSE 3001
# Run the hello binary.
ENTRYPOINT ["/freezeframe"]
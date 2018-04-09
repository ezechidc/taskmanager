# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.9 AS builder

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/ezechidc/taskmanager

# Setting up working directory
WORKDIR /go/src/github.com/ezechidc/taskmanager

# Get dep for managing and restoring dependencies
RUN go get -u github.com/golang/dep/cmd/dep


# copies the Gopkg.toml and Gopkg.lock to WORKDIR
COPY Gopkg.toml Gopkg.lock ./

# Restore dep dependencies
RUN dep ensure -vendor-only

# Build the taskmanager command inside the container.
RUN go install github.com/ezechidc/taskmanager

# Run the taskmanager command when the container starts.
ENTRYPOINT /go/bin/taskmanager

# Service listens on port 8080.
EXPOSE 8080
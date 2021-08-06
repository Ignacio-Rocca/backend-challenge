FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/golang/checkout_service

RUN go install github.com/golang/checkout_service

# Run the  command by default when the container starts.
ENTRYPOINT /go/bin/outyet

# Document that the service listens on port 8080.
EXPOSE 8081
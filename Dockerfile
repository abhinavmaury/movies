FROM golang:alpine as builder

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/abhinavmaury/movies

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

RUN go build .

# Run the executable
CMD ["movies"]
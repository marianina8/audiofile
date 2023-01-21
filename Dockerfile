FROM golang:1.19

# Set the working directory
WORKDIR /audiofile

# Copy the go.mod and go.sum files
COPY go.mod go.sum main.go Makefile ./

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN make build-all

# Run the tests
RUN make test-verbose

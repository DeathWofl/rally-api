FROM golang:1.14-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-sample-app .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/go-sample-app"]

# FROM golang:1.14

# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

# CMD ["app"]

# FROM golang:1.14

# RUN mkdir /app

# ADD . /app

# WORKDIR /app
# COPY . /go/src/app

# # RUN go get -d ./go/src/app
# RUN go build -o wine-reviews

# CMD ["/go/src/app/wine-reviews"]
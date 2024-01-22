# Stage 1: Build stage
FROM golang:1.21-alpine3.18  as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# ADD . /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY *.go ./

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "./cmd/main.go"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -o bin/go-service main .

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main . 
# copy env file to root directory 
COPY --from=builder /app/.env .

# Command to run the executable
CMD ["./main"]
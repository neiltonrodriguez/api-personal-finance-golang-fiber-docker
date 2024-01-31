FROM golang:1.21-alpine3.18  as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main . 

COPY --from=builder /app/.env .

CMD ["./main"]
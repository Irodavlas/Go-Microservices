FROM golang:1.21.3 as builder
WORKDIR /app/
COPY user-service/go.build.mod ./go.mod
COPY user-service/go.sum ./

# Copy the common/response directory into the container
COPY ./common/response/ ./common/response/

RUN go mod download
COPY ./user-service .

RUN CGO_ENABLED=0 GOOS=linux go build -o service .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/service . 
EXPOSE 5000

CMD ["./service"]
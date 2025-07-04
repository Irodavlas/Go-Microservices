FROM golang:1.21.3 AS builder
WORKDIR /app/

COPY api-gateway/go.build.mod ./go.mod
COPY api-gateway/go.sum ./

# COPY response module into the container but at the root because
# it's being accessed from the go.build.mod file as ./common/response 
# and the context is withing the root dir website/

COPY ./common/response ../common/response 
RUN go mod download

COPY ./api-gateway .

RUN CGO_ENABLED=0 GOOS=linux go build -o service

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/service .

EXPOSE 8000

CMD [ "./service" ]
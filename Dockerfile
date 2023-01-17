FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/Swapica/rpc-proxy-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/rpc-proxy-svc /go/src/github.com/Swapica/rpc-proxy-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/rpc-proxy-svc /usr/local/bin/rpc-proxy-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["rpc-proxy-svc"]

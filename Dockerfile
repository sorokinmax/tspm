##
## Build stage
##
FROM golang:1.19.2-alpine3.16 AS builder
RUN apk update && apk upgrade
WORKDIR /src
COPY . .

RUN go mod download
RUN GOOS=linux go build -o /tspm .

##
## Final image stage
##
FROM alpine:3.16
WORKDIR /
COPY --from=builder /tspm /tspm

ENTRYPOINT ["/tspm"]
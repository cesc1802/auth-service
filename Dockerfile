#build stage
FROM golang:1.18-buster AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app

#final stage
FROM ubuntu:20.04
WORKDIR /
COPY messages messages
COPY migrations migrations
COPY --from=builder /go/bin/app /app
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Etc/UTC
RUN apt-get update -y
RUN apt-get -y install tzdata
ENTRYPOINT ["/app"]
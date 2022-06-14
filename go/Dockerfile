FROM golang:latest AS go_base
COPY ./apiserver /go
WORKDIR /go
ENV GOPATH /go/bin
RUN mkdir -p /go/bin
RUN go mod download
RUN go build -o /go/bin /go/main.go
ENTRYPOINT ["/go/bin/main"]

FROM alpine:latest
RUN apk add libc6-compat
WORKDIR /apiserver
COPY --from=go_base /go/bin/main ./main

ENTRYPOINT ["./main"]
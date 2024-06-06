FROM golang:1.21 as Build

WORKDIR /go/src/stress-tool

RUN apt update

COPY . .

RUN go get -d -v ./... && \
    go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o /go/bin/stress-tool ./cmd/stress-tool

FROM alpine:3.18

COPY --from=Build /go/bin/stress-tool /usr/local/bin/stress-tool

ENTRYPOINT ["/usr/local/bin/stress-tool"]

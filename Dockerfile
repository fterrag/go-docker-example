FROM golang:1.10.1-alpine as builder
WORKDIR /go/src/github.com/fterrag/go-docker-example
COPY . .
RUN apk --no-cache add git
RUN go get -u github.com/golang/dep/cmd/dep \
    && go get golang.org/x/tools/cmd/cover \
    && go get github.com/mattn/goveralls
RUN dep ensure
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/fterrag/go-docker-example/app .
CMD ["./app"]

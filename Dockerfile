FROM golang:1.9.2-alpine as builder
WORKDIR /go/src/github.com/fterrag/go-docker-example
COPY . .
RUN go get -d -v .
RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/fterrag/go-docker-example/app .
CMD ["./app"]

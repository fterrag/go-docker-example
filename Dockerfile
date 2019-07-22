FROM golang:1.12-alpine as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get \
    && go get golang.org/x/tools/cmd/cover \
    && go get github.com/mattn/goveralls

RUN go build -o example cmd/example/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/example .
CMD ["./example"]

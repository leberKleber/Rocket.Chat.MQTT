FROM golang:1.11 as buildContainer

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOPATH=/

COPY . /src/rocket.chat.mqtt
WORKDIR /src/rocket.chat.mqtt

RUN go get ./... &&\
    go build -ldflags -s -a -installsuffix cgo -o rocket.chat.mqtt ./cmd/mqtt/


FROM alpine

COPY --from=buildContainer /src/rocket.chat.mqtt/rocket.chat.mqtt /rocket.chat.mqtt

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ENTRYPOINT ["/rocket.chat.mqtt"]
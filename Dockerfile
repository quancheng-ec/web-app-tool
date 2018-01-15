FROM golang:1.9 as builder

RUN go get -t github.com/urfave/cli && \
    go get -t github.com/visionmedia/go-debug && \
    go get -t github.com/kelseyhightower/envconfig && \
    go get -t github.com/denverdino/aliyungo/oss && \
    go get -t github.com/codeskyblue/go-sh

ADD . $GOPATH/src/web-app-tool

RUN go build $GOPATH/src/web-app-tool/web-app-tool.go


FROM node:alpine

WORKDIR /root/

RUN npm install cnpm -g

COPY --from=builder /go/src/web-app-tool/web-app-tool .

CMD ["./web-app-tool"]


FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && apk --no-cache add git

WORKDIR /app/gomicro-server

COPY . .    

RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gomicro-server

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/gomicro-server/gomicro-server .

CMD [ "./gomicro-server" ]

FROM golang:1.21-alpine AS build

ENV TZ "Asia/Shanghai"

RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --no-cache add ca-certificates \
  && update-ca-certificates

WORKDIR /app

ENV GOPROXY=https://goproxy.cn/,direct GO111MODULE=on

COPY ./go.mod ./
# COPY ./go.sum ./

RUN go mod tidy

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway ./

FROM alpine

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=build /app/gateway /app/gateway
COPY --from=build /app/config /app/config
RUN chmod +x /app/gateway

EXPOSE 8081

CMD ["/app/gateway", "server"]

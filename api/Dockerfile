FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/bookstore api/bookstore.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Taipei /usr/share/zoneinfo/Asia/Taipei
ENV TZ Asia/Taipei

WORKDIR /app
COPY --from=builder /app/bookstore /app/bookstore
COPY --from=builder /app/etc /app/etc

CMD ["./bookstore", "-f", "etc/bookstore-api.yaml"]

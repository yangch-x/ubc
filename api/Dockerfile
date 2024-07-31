FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装必要的包，包括 tzdata 和 Python
RUN apk update --no-cache && apk add --no-cache tzdata python3 py3-pip

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY api/etc/ /app/etc/
COPY py/ /app/py/
RUN go build -ldflags="-s -w" -o /app/backpack api/backpack.go

# 使用一个较小的基础镜像来运行最终的应用
FROM alpine:latest

# 安装必要的运行时依赖，包括 Python 和 pip
RUN apk add --no-cache \
    ca-certificates \
    python3 \
    py3-pip \
    && pip3 install --no-cache-dir openai PyPDF2 termcolor

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/backpack /app/backpack
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/py /app/py

EXPOSE 8888

CMD ["./backpack", "-f", "etc/backpack.yaml"]

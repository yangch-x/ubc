# 使用 Golang 的基础镜像来构建应用
FROM golang:1.19-alpine3.18 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装必要的包，包括 tzdata 和 Python
RUN apk update --no-cache && apk add --no-cache tzdata python3 py3-pip

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY api/etc/ /app/etc/
COPY py/ /app/py/
RUN go build -ldflags="-s -w" -o /app/backpack api/backpack.go

# 使用一个较小的基础镜像来运行最终的应用
FROM alpine:3.18

# 安装必要的运行时依赖，包括 Python 和 pip
RUN apk add --no-cache \
    ca-certificates \
    python3 \
    py3-pip \
    curl

# 创建虚拟环境并安装 Python 包
RUN python3 -m venv /venv \
    && /venv/bin/pip install --no-cache-dir \
        openai \
        PyPDF2 \
        termcolor \
        reportlab \
        requests

# 创建非root用户
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/backpack /app/backpack
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/py /app/py

# 更改文件所有权
RUN chown -R appuser:appgroup /app

# 设置环境变量以使用虚拟环境中的 Python
ENV PATH="/venv/bin:$PATH"

# 切换到非root用户
USER appuser

EXPOSE 8888

# 添加健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8888/ping || exit 1

CMD ["./backpack", "-f", "etc/backpack.yaml"]

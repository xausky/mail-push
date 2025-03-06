# 使用官方 Go 镜像作为构建环境
FROM golang:alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 go build -o mail-push

# 使用轻量级的 alpine 作为运行环境
FROM alpine:latest

# 安装 ca-certificates，用于 SMTP SSL 连接
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从构建阶段复制二进制文件和必要的文件
COPY --from=builder /app/mail-push .
COPY --from=builder /app/config.toml .
COPY --from=builder /app/static ./static

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./mail-push"] 
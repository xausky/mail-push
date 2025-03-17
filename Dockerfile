# 前端构建阶段
FROM node:18-alpine AS frontend-builder

# 设置工作目录
WORKDIR /app

# 复制前端项目文件
COPY . ./
RUN npm install

# 构建前端项目
RUN npm run build

# 后端构建阶段
FROM golang:alpine AS backend-builder

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

# 最终运行阶段
FROM alpine:latest

# 安装 ca-certificates，用于 SMTP SSL 连接
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从构建阶段复制二进制文件和必要的文件
COPY --from=backend-builder /app/mail-push .
COPY --from=backend-builder /app/config.toml .
COPY --from=frontend-builder /app/build ./build

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./mail-push"] 
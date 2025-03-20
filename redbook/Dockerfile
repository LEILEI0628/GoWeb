# 基础镜像
FROM ubuntu:20.04

COPY redbook /app/redbook

# 设置工作目录
WORKDIR /app

ENTRYPOINT ["/app/redbook"]

# Build命令：docker build -t redbook:v0.0.1 .
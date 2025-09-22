FROM golang:1.24.4-alpine AS builder
# 国内代理
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine
COPY --from=builder /app/main /main

# 🔥 添加这一行 - 告诉容器启动时运行什么
CMD ["/main"]
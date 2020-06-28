# Go实现网关服务
> 框架采用GoFrame框架

### 安装OpenTracing链路追踪
```bash
docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

### 生成swagger文档
1. [安装gf-cli](http://127.0.0.1:8100/docs/swagger.json)
2. 执行 gf swagger
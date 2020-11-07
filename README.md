# Go实现网关服务
> 框架采用GoFrame框架

## 已实现功能
- [x] http/tcp/grpc 代理
- [x] 四种策略的负载均衡
- [x] 接入swagger
- [x] 异常code错误码
- [x] jwt鉴权token
- [x] 基础链路追踪中间件


## Tips
1. 安装OpenTracing链路追踪
```bash
docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

2. 生成swagger文档
1. [安装gf-cli](http://127.0.0.1:8100/docs/swagger.json)
2. 执行 gf swagger 或者 go generate
3. 访问 http://127.0.0.1:8100/swagger 即可

## TODO
- [ ] Mysql加入链路追踪机制
- [ ] 加入redis操作并新增链路追踪
- [ ] 加入限流熔断中间件（Sentinel + Nacos）
- [ ] 新增Prometheus埋点
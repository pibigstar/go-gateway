# 代理

## 正向代理(forward-proxy)
>web浏览器请求数据时先经过我们得正向代理服务，然后再将请求发送到对应的服务器

1. 计算机设置web代理为: 127.0.0.1:9000
![](../screenshot/setting.png)
2. 启动正向代理（main.go）
3. web浏览器访问http地址
> 必须是http,这里只代理了http服务，没有代理https



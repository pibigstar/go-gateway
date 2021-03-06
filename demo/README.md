# 代理

## 1. 正向代理(forward-proxy)
> 正向代理作用于客户端，它可以帮助客户端来隐藏客户的真实IP或一些需要特殊环境才能访问的服务（翻墙）
> web浏览器请求数据时先经过我们得正向代理服务，然后再将请求发送到对应的服务器

1. 计算机设置web代理为: 127.0.0.1:9000
![](../screenshot/setting.png)
2. 启动正向代理（main.go）
3. web浏览器访问http地址
> 必须是http,这里只代理了http服务，没有代理https

## 2. 反向代理（reverse）
> 反向代理作用于服务器端，主要目的是为了隐藏服务器端真实地址，做负载均衡，以及流量控制等。
>
> 当我们访问 127.0.0.1:7000 端口时，其实这个端口运行的是一个反向代理服务器，
> 它会将请求发送到该反向代理服务器代理的地址: 127.0.0.1:7001 上

### 2.1 最简单的反向代理（base）
> 只是简单的更改一下HOST地址，实现一个简单的反向代理

### 2.2 高级点反向代理（http）
> 使用`httputil`包提供的ReverseProxy, 它可以很方便的帮我们做反向代理地址的处理，日志记录，返回内容修改等。


## 3. 负载均衡（balance）

### 3.1 随机负载均衡（random）
> 从目标主机数组中随机获取一个目标主机返回

### 3.2 轮询负载均衡（polling)
> 从目标数组中获取当前索引下对应的目标主机，然后将索引值加一再对目标长度取余即可。

### 3.3 加权轮询负载均衡（weight_polling)
> 权重越大，出现的概率越大


## 4. 中间件（middleware)
> 通过将中间件函数放到一个切片中，依次执行，最后执行到业务函数为止
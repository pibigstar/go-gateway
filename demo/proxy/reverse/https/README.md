#  生成证书与私钥

证书签名生成方式:
```bash
# CA私钥
openssl genrsa -out ca.key 2048

# CA数据证书
openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 5000 -out ca.crt

# 服务器私钥（默认由CA签发）
openssl genrsa -out server.key 2048

# 服务器证书签名请求：Certificate Sign Request，简称csr（localhost代表你的域名）
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr

# 上面2个文件生成服务器证书（days代表有效期）
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
```

注意： 如果谷歌浏览器认为该地址是不安全的不让连接，直接 在该页面 键盘输入 `thisisunsafe` 即可。